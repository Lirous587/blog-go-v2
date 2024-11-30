package repository

import (
	"blog/models"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
	"sync"
)

const (
	invalidLabelIds = "无效的标签"
)

type EssayRepo interface {
	Create(data *models.EssayParams) error
	Read(id int) (*models.EssayData, error)
	Update(data *models.EssayUpdateParams) error
	Delete(id int) error
	GetList(query *models.EssayQuery) (*models.EssayListAndPage, error)
	GetRecommendList() ([]models.EssayData, error)
	GetAll() ([]models.EssayData, error)
}

type essayRawData struct {
	models.EssayData
	LabelIDs          string `db:"label_ids"`
	LabelNames        string `db:"label_names"`
	LabelIntroduction string `db:"label_introduction"`
}

type EssayRepoMySQL struct {
	db *sqlx.DB
}

func NewEssayRepoMySQL(db *sqlx.DB) *EssayRepoMySQL {
	return &EssayRepoMySQL{
		db: db,
	}
}

func (r *EssayRepoMySQL) Create(data *models.EssayParams) error {
	return newSqlxTx(r.db, func(tx *sqlx.Tx) error {
		result, err := r.insertEssay(tx, data)
		if err != nil {
			return fmt.Errorf("r.insert(tx, data) failed: %w", err)
		}
		eid64, err := result.LastInsertId()
		if err != nil {
			return err
		}
		eid := int(eid64)
		if err := r.insertLabels(tx, eid, data.LabelIds); err != nil {
			return fmt.Errorf("r.insertLabels(tx, eid, data.LabelIds) failed: %w", err)
		}
		return nil
	})
}

func (r *EssayRepoMySQL) insertEssay(tx *sqlx.Tx, p *models.EssayParams) (sql.Result, error) {
	sqlStr := `
        INSERT INTO essay(name, kind_id, if_top, content, if_recommend, introduction, created_time, img_id,keywords) 
        VALUES (:name, :kind_id, :if_top, :content, :if_recommend, :introduction, :created_time,:img_id,:keywords)
    `
	result, err := tx.NamedExec(sqlStr, p)

	return result, err
}

func (r *EssayRepoMySQL) Read(id int) (*models.EssayData, error) {
	var wg sync.WaitGroup
	taskCount := 4
	wg.Add(taskCount)
	data := new(models.EssayData)
	var errChan = make(chan error, taskCount)

	data.ID = id
	go func() {
		defer wg.Done()
		if err := r.readEssay(data); err != nil {
			errChan <- fmt.Errorf("r.readEssay(data) failed,err:%w", err)
			return
		}
	}()
	go func() {
		defer wg.Done()
		labelList, err := r.readLabels(data.ID)
		if err != nil {
			errChan <- fmt.Errorf("r.readLabels(data.ID) failed,err:%w", err)
			return
		}
		data.LabelList = labelList
	}()
	go func() {
		defer wg.Done()
		if err := r.addVisitedTimes(data.ID); err != nil {
			errChan <- fmt.Errorf("r.addVisitedTimes(data.ID) failed,err:%w", err)
			return
		}
	}()
	go func() {
		defer wg.Done()
		nearByEssay := make([]models.EssayData, 0, 5)
		if err := r.getNearbyEssays(&nearByEssay, data.KindID, data.ID); err != nil {
			errChan <- fmt.Errorf("increaseEssayCount failed,err:%w", err)
			return
		}
		data.NearEssayList = nearByEssay
	}()

	wg.Wait()
	close(errChan)
	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func (r *EssayRepoMySQL) readEssay(data *models.EssayData) error {
	sqlStr := `
		SELECT e.id,e.name,e.kind_id, e.content, e.introduction,e.img_id, e.created_time, e.visited_times, e.keywords,g.img_url,
			k.name AS kind_name
		FROM essay e
		LEFT JOIN blog.gallery g on g.id = e.img_id
		LEFT JOIN e_kind k on e.kind_id = k.id
		where e.id = ?
		`
	return r.db.Get(data, sqlStr, data.ID)
}

func (r *EssayRepoMySQL) addVisitedTimes(id int) error {
	sqlStr := `
	UPDATE essay SET visited_times = visited_times + 1
		WHERE id = ?`
	_, err := r.db.Exec(sqlStr, id)
	return err
}

func (r *EssayRepoMySQL) getNearbyEssays(data *[]models.EssayData, kID int, eID int) error {
	sqlStr := `
		(SELECT e.id, e.name, e.kind_id, e.introduction, e.created_time, g.img_url,e.visited_times,
		 	k.name AS kind_name
			FROM essay e
			LEFT JOIN e_kind k on k.id = e.kind_id
			LEFT JOIN blog.gallery g on g.id = e.img_id
			WHERE e.kind_id = ? AND e.id < ?
			ORDER BY e.id
			LIMIT 2)
		UNION ALL
		(SELECT e.id, e.name, e.kind_id, e.introduction, e.created_time,g.img_url,e.visited_times,
		 	k.name AS kindName
			FROM essay e
			LEFT JOIN e_kind k on k.id = e.kind_id
			LEFT JOIN blog.gallery g on g.id = e.img_id
			WHERE e.kind_id = ? AND e.id > ?
			ORDER BY e.id
		LIMIT 2)
  `
	return r.db.Select(data, sqlStr, kID, eID, kID, eID)
}

func (r *EssayRepoMySQL) Update(data *models.EssayUpdateParams) error {
	return newSqlxTx(r.db, func(tx *sqlx.Tx) error {
		// 更新essay表
		if err := r.updateEssay(tx, data); err != nil {
			return fmt.Errorf("r.updateEssay(tx, data) failed,err%w", err)
		}
		// 更新essay_label
		if err := r.updateLabel(tx, data.ID, data.OldLabelIds, data.LabelIds); err != nil {
			return fmt.Errorf("r.updateLabel(tx, data) failed,err:%w", err)
		}
		return nil
	})
}

func (r *EssayRepoMySQL) updateEssay(tx *sqlx.Tx, data *models.EssayUpdateParams) error {
	sqlStr := `UPDATE essay SET 
               name = :name,
               kind_id = :kind_id,
               introduction = :introduction,
               content = :content,
               img_id = :img_id,
               if_top = :if_top,
               if_recommend = :if_recommend,
               keywords = :keywords
               WHERE id = :id`
	_, err := tx.NamedExec(sqlStr, data)
	return err
}

// Delete essay中删除数据 e_label中删除数据
func (r *EssayRepoMySQL) Delete(id int) (err error) {
	return newSqlxTx(r.db, func(tx *sqlx.Tx) error {
		if err := r.deleteLabels(tx, id); err != nil {
			return fmt.Errorf("r.deleteLabels(tx, id) failed,err:%w", err)
		}
		err := r.deleteEssay(tx, id)
		if err != nil {
			return fmt.Errorf("r.deleteEssay(tx, id) failed,err:%w", err)
		}
		return nil
	})
}

func (r *EssayRepoMySQL) deleteEssay(tx *sqlx.Tx, eid int) (err error) {
	sqlStr := `DELETE FROM essay WHERE id = ?`
	_, err = tx.Exec(sqlStr, eid)
	return err
}

func (r *EssayRepoMySQL) GetAll() (list []models.EssayData, err error) {
	sqlStr := `
		SELECT e.id, e.name, e.created_time,e.visited_times,e.kind_id,e.introduction,e.if_top,e.if_recommend, e.img_id,g.img_url,
		       k.name AS kind_name,
		    COALESCE(GROUP_CONCAT(el.label_id), '') AS label_ids,
            COALESCE(GROUP_CONCAT(l.name), '') AS label_names
		FROM essay e
		LEFT JOIN e_kind k ON e.kind_id = k.id
		LEFT JOIN eid_lid el ON e.id = el.essay_id
		LEFT JOIN e_label l ON l.id = el.label_id
		LEFT JOIN gallery g on g.id = e.img_id	
		GROUP BY e.id,g.img_url
		ORDER BY e.id DESC
	`
	var rawDataList = new([]essayRawData)
	if err = r.db.Select(rawDataList, sqlStr); err != nil {
		return
	}
	list = make([]models.EssayData, len(*rawDataList))
	for i, raw := range *rawDataList {
		list[i] = raw.EssayData
		if raw.LabelNames != "" && raw.LabelIDs != "" {
			ids := strings.Split(raw.LabelIDs, ",")
			names := strings.Split(raw.LabelNames, ",")
			list[i].LabelList = make([]models.EssayLabelData, len(ids))
			for j := range ids {
				id, _ := strconv.Atoi(ids[j])
				list[i].LabelList[j] = models.EssayLabelData{
					ID:   id,
					Name: names[j],
				}
			}
		}
	}
	return
}

func (r *EssayRepoMySQL) GetRecommendList() (list []models.EssayData, err error) {
	list = make([]models.EssayData, 0, 5)
	sqlStr := `
		SELECT e.id, e.name, e.created_time, g.img_url 
		FROM essay e 
		JOIN blog.gallery g on e.img_id = g.id
		WHERE if_recommend = true
		ORDER BY e.id DESC 
		LIMIT 5
	`
	err = r.db.Select(&list, sqlStr)
	return
}

func (r *EssayRepoMySQL) GetList(query *models.EssayQuery) (*models.EssayListAndPage, error) {
	var wg sync.WaitGroup
	taskCount := 2
	var errChan = make(chan error, taskCount)
	wg.Add(taskCount)

	data := new(models.EssayListAndPage)

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
		list, err := r.getEssayList(whereClause, args)
		if err != nil {
			errChan <- fmt.Errorf("getList failed,err:%w", err)
			return
		}
		data.EssayList = list
	}()

	go func() {
		defer wg.Done()
		totalCount, err := r.getEssayCount(whereClause, args)
		if err != nil {
			errChan <- fmt.Errorf("getList failed,err:%w", err)
			return
		}
		data.TotalPages = (totalCount + query.PageSize - 1) / query.PageSize
	}()

	wg.Wait()

	close(errChan)

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (r *EssayRepoMySQL) getEssayList(whereClause string, args []interface{}) (list []models.EssayData, err error) {
	list = make([]models.EssayData, 0, 5)
	rawDataList := make([]essayRawData, 0, 5)
	baseSelect := `
     SELECT e.id, e.name, e.kind_id, e.if_recommend, e.if_top, e.introduction,e.keywords,
        e.created_time, e.visited_times,e.img_id,g.img_url,
        k.name AS kind_name,
        COALESCE(GROUP_CONCAT(el.label_id), '') AS label_ids,
        COALESCE(GROUP_CONCAT(l.name), '') AS label_names,
        COALESCE(GROUP_CONCAT(l.introduction), '') AS label_introduction
    FROM essay e
    LEFT JOIN gallery g on g.id = e.img_id
    LEFT JOIN e_kind k ON e.kind_id = k.id
    LEFT JOIN eid_lid el ON e.id = el.essay_id
    LEFT JOIN e_label l ON l.id = el.label_id  
	`
	groupBy := "GROUP BY e.id,g.img_url"

	// 优先排序ifTop为true的记录
	orderBy := "ORDER BY e.if_top DESC, e.id DESC"

	sqlStr := fmt.Sprintf("%s %s %s %s LIMIT ? OFFSET ?",
		baseSelect, whereClause, groupBy, orderBy)

	if err := r.db.Select(&rawDataList, sqlStr, args...); err != nil {
		return nil, err
	}

	// 处理查询结果
	list = make([]models.EssayData, len(rawDataList))
	for i, raw := range rawDataList {
		list[i] = raw.EssayData
		// 处理标签数据
		if raw.LabelIDs != "" && raw.LabelNames != "" {
			ids := strings.Split(raw.LabelIDs, ",")
			names := strings.Split(raw.LabelNames, ",")
			introduction := strings.Split(raw.LabelIntroduction, ",")
			list[i].LabelList = make([]models.EssayLabelData, len(ids))

			for j := range ids {
				id, _ := strconv.Atoi(ids[j])
				list[i].LabelList[j] = models.EssayLabelData{
					ID:           id,
					Name:         names[j],
					Introduction: introduction[j],
				}
			}
		}
	}
	return list, nil
}

func (r *EssayRepoMySQL) getEssayCount(whereClause string, args []interface{}) (totalCount int, err error) {
	baseSql := `
        SELECT COUNT(DISTINCT e.id)
        FROM essay e 
        LEFT JOIN e_kind k ON e.kind_id = k.id
        LEFT JOIN eid_lid el ON e.id = el.essay_id
   	`

	sqlStr := fmt.Sprintf("%s %s", baseSql, whereClause)
	if err = r.db.Get(&totalCount, sqlStr, args[:len(args)-2]...); err != nil {
		return
	}
	return
}
