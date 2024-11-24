package mysql

import (
	"blog/models"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type essayRawData struct {
	models.EssayData
	LabelIDs          string `db:"label_ids"`
	LabelNames        string `db:"label_names"`
	LabelIntroduction string `db:"label_introduction"`
}

func GetRecommendEssayList(data *[]models.EssayData) error {
	sqlStr := `
		SELECT e.id, e.name, e.created_time, g.img_url 
		FROM essay e 
		JOIN blog.gallery g on e.img_id = g.id
		WHERE if_recommend = true
		ORDER BY e.id DESC 
		LIMIT 5
	`
	return db.Select(data, sqlStr)
}

func GetEssayList(data *models.EssayListAndPage, query models.EssayQuery) error {
	var wg sync.WaitGroup
	taskCount := 2
	var errChan = make(chan error, taskCount)
	wg.Add(taskCount)

	// 计算偏移量
	offset := (query.Page - 1) * query.PageSize
	where := make([]string, 0)
	args := make([]interface{}, 0)

	if query.LabelID != 0 {
		where = append(where, "el.label_id = ?")
		args = append(args, query.LabelID)
	}
	if query.KindID != 0 {
		where = append(where, "e.kind_id = ?")
		args = append(args, query.KindID)
	}

	whereClause := ""
	if len(where) > 0 {
		whereClause = "WHERE " + strings.Join(where, " AND ")
	}

	args = append(args, query.PageSize, offset)

	go func() {
		defer wg.Done()
		if err := getEssayList(data, whereClause, args); err != nil {
			errChan <- fmt.Errorf("getList failed,err:%w", err)
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := getEssayCount(data, query.PageSize, whereClause, args); err != nil {
			errChan <- fmt.Errorf("getList failed,err:%w", err)
			return
		}
	}()

	wg.Wait()

	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func getEssayList(data *models.EssayListAndPage, whereClause string, args []interface{}) error {
	rawDataList := make([]essayRawData, 0, 5)
	baseSelect := `
     SELECT e.id, e.name, e.kind_id, e.if_recommend, e.if_top, e.introduction,
        e.created_time, e.visited_times,e.img_id,g.img_url,
        k.name AS kind_name,
        COALESCE(GROUP_CONCAT(el.label_id), '') AS label_ids,
        COALESCE(GROUP_CONCAT(l.name), '') AS label_names,
        COALESCE(GROUP_CONCAT(l.introduction), '') AS label_introduction
    FROM essay e
    LEFT JOIN gallery g on g.id = e.img_id
    LEFT JOIN kind k ON e.kind_id = k.id
    LEFT JOIN essay_label el ON e.id = el.essay_id
    LEFT JOIN label l ON l.id = el.label_id  
	`

	groupBy := "GROUP BY e.id,g.img_url"

	// 优先排序ifTop为true的记录
	orderBy := "ORDER BY e.if_top DESC, e.id DESC"

	sqlStr := fmt.Sprintf("%s %s %s %s LIMIT ? OFFSET ?",
		baseSelect, whereClause, groupBy, orderBy)

	if err := db.Select(&rawDataList, sqlStr, args...); err != nil {
		return err
	}

	// 处理查询结果
	data.EssayList = make([]models.EssayData, len(rawDataList))
	for i, raw := range rawDataList {
		data.EssayList[i] = raw.EssayData
		// 处理标签数据
		if raw.LabelIDs != "" && raw.LabelNames != "" {
			ids := strings.Split(raw.LabelIDs, ",")
			names := strings.Split(raw.LabelNames, ",")
			introduction := strings.Split(raw.LabelIntroduction, ",")
			data.EssayList[i].LabelList = make([]models.LabelData, len(ids))

			for j := range ids {
				id, _ := strconv.Atoi(ids[j])
				data.EssayList[i].LabelList[j] = models.LabelData{
					ID:           id,
					Name:         names[j],
					Introduction: introduction[j],
				}
			}
		}
	}
	return nil
}

func getEssayCount(data *models.EssayListAndPage, PageSize int, whereClause string, args []interface{}) error {
	baseSql := `
        SELECT COUNT(DISTINCT e.id)
        FROM essay e 
        LEFT JOIN kind k ON e.kind_id = k.id
        LEFT JOIN essay_label el ON e.id = el.essay_id
   	`

	var totalCount int
	sqlStr := fmt.Sprintf("%s %s", baseSql, whereClause)
	if err := db.Get(&totalCount, sqlStr, args[:len(args)-2]...); err != nil {
		return err
	}

	data.TotalPages = (totalCount + PageSize - 1) / PageSize
	return nil
}

func GetAllEssay(data *[]models.EssayData) error {
	sqlStr := `
		SELECT e.id, e.name, e.created_time,e.visited_times,e.kind_id,e.introduction,e.if_top,e.if_recommend, e.img_id,g.img_url,
		       k.name AS kind_name,
		    COALESCE(GROUP_CONCAT(el.label_id), '') AS label_ids,
            COALESCE(GROUP_CONCAT(l.name), '') AS label_names
		FROM essay e
		LEFT JOIN kind k ON e.kind_id = k.id
		LEFT JOIN essay_label el ON e.id = el.essay_id
		LEFT JOIN label l ON l.id = el.label_id
		LEFT JOIN blog.gallery g on g.id = e.img_id	
		GROUP BY e.id,g.img_url
		ORDER BY e.id DESC
	`
	var err error
	var rawDataList = new([]essayRawData)
	if err = db.Select(rawDataList, sqlStr); err != nil {
		return err
	}
	*data = make([]models.EssayData, len(*rawDataList))
	for i, raw := range *rawDataList {
		(*data)[i] = raw.EssayData
		if raw.LabelNames != "" && raw.LabelIDs != "" {
			ids := strings.Split(raw.LabelIDs, ",")
			names := strings.Split(raw.LabelNames, ",")
			(*data)[i].LabelList = make([]models.LabelData, len(ids))
			for j := range ids {
				id, _ := strconv.Atoi(ids[j])
				(*data)[i].LabelList[j] = models.LabelData{
					ID:   id,
					Name: names[j],
				}
			}
		}
	}
	return err
}
