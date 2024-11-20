package mysql

import (
	"blog/models"
	"blog/utils"
	"fmt"
	"strings"
	"sync"
)

func GetGalleryList(data *models.GalleryListAndPage, query models.GalleryQuery) error {
	var wg sync.WaitGroup
	taskCount := 2
	var errChan = make(chan error, taskCount)
	wg.Add(taskCount)

	// 计算偏移量
	offset := (query.Page - 1) * query.PageSize
	args := make([]interface{}, 0)

	where := make([]string, 0)
	if query.KindID != 0 {
		where = append(where, "g.kind_id = ?")
		args = append(args, query.KindID)
	}

	//为后续查询筛选条件做准备
	whereClause := ""
	if len(where) > 0 {
		whereClause = "WHERE " + strings.Join(where, " AND ")
	}

	args = append(args, query.PageSize, offset)

	go func() {
		defer wg.Done()
		if err := getGalleryList(data, whereClause, args); err != nil {
			errChan <- fmt.Errorf("getList failed,err:%w", err)
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := getGalleryCount(data, query.PageSize, whereClause, args); err != nil {
			errChan <- fmt.Errorf("getCount failed,err:%w", err)
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

func getGalleryList(data *models.GalleryListAndPage, whereClause string, args []interface{}) error {
	rawDataList := make([]models.GalleryData, 0, 10)
	baseSelect := `
        SELECT g.id,g.img_url,g.kind_id,k.name
        	FROM gallery g
        	LEFT JOIN gallery_kind k ON g.kind_id = k.id
        	`

	orderBy := "ORDER BY g.id DESC"

	sqlStr := fmt.Sprintf("%s %s %s  LIMIT ? OFFSET ?",
		baseSelect, whereClause, orderBy)

	if err := db.Select(&rawDataList, sqlStr, args...); err != nil {
		return err
	}

	// 处理查询结果
	data.GalleryList = make([]models.GalleryData, len(rawDataList))
	for i, raw := range rawDataList {
		data.GalleryList[i] = raw
	}
	return nil
}

func getGalleryCount(data *models.GalleryListAndPage, PageSize int, whereClause string, args []interface{}) error {
	baseSql := `
        SELECT COUNT(DISTINCT g.id)
        FROM gallery g  
   	`

	var totalCount int
	sqlStr := fmt.Sprintf("%s %s", baseSql, whereClause)
	if err := db.Get(&totalCount, sqlStr, args[:len(args)-2]...); err != nil {
		return err
	}

	data.TotalPage = (totalCount + PageSize - 1) / PageSize
	return nil
}

func CreateGallery(p *models.GalleryParams) error {
	p.Url = utils.SanitizedFileName(p.Url)
	sqlStr := `INSERT INTO gallery(img_url, kind_id) VALUES (:img_url,:kind_id)`
	_, err := db.NamedExec(sqlStr, p)
	return err
}

func DeleteGallery(id int) error {
	sqlStr := `DELETE FROM gallery WHERE id = ?`
	_, err := db.Exec(sqlStr, id)
	return err
}

func UpdateGallery(p *models.GalleryUpdateParams) error {
	sqlStr := `UPDATE gallery SET img_url = :img_url,kind_id = :kind_id WHERE id = :id`
	_, err := db.NamedExec(sqlStr, p)
	return err
}
