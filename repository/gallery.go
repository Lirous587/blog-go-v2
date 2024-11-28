package repository

import (
	"blog/models"
	"blog/utils"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"sync"
)

type GalleryRepo interface {
	Create(data *models.GalleryParams) error
	Read(id int) (*models.GalleryData, error)
	Update(data *models.GalleryUpdateParams) error
	Delete(id int) error
	GetList(query *models.GalleryQuery) (*models.GalleryListAndPage, error)
	getList(data *models.GalleryListAndPage, whereClause string, args []interface{}) error
	getCount(data *models.GalleryListAndPage, PageSize int, whereClause string, args []interface{}) error
}

type GalleryRepoMySQL struct {
	db *sqlx.DB
}

func NewGalleryRepoMySQL(db *sqlx.DB) *GalleryRepoMySQL {
	return &GalleryRepoMySQL{
		db: db,
	}
}

func (r *GalleryRepoMySQL) Create(data *models.GalleryParams) error {
	data.ImgUrl = utils.SanitizedFileName(data.ImgUrl)
	sqlStr := `INSERT INTO gallery(img_url, kind_id) VALUES (:img_url,:kind_id)`
	_, err := r.db.NamedExec(sqlStr, data)
	return err
}

func (r *GalleryRepoMySQL) Read(id int) (data *models.GalleryData, err error) {
	sqlStr := `SELECT content,source,img_id,if_could_type FROM heart_words WHERE id = ?`
	err = r.db.Get(data, sqlStr, id)
	return data, err
}

func (r *GalleryRepoMySQL) Update(data *models.GalleryUpdateParams) error {
	sqlStr := `UPDATE gallery SET img_url = :img_url,kind_id = :kind_id WHERE id = :id`
	_, err := r.db.NamedExec(sqlStr, data)
	return err
}

func (r *GalleryRepoMySQL) Delete(id int) error {
	sqlStr := `DELETE FROM gallery WHERE id = ?`
	_, err := r.db.Exec(sqlStr, id)
	return err
}

func (r *GalleryRepoMySQL) GetList(query *models.GalleryQuery) (data *models.GalleryListAndPage, err error) {
	data = new(models.GalleryListAndPage)
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
		if err := r.getList(data, whereClause, args); err != nil {
			errChan <- fmt.Errorf("getList failed,err:%w", err)
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := r.getCount(data, query.PageSize, whereClause, args); err != nil {
			errChan <- fmt.Errorf("getCount failed,err:%w", err)
			return
		}
	}()

	wg.Wait()

	close(errChan)

	for err = range errChan {
		if err != nil {
			return
		}
	}
	return
}

func (r *GalleryRepoMySQL) getList(data *models.GalleryListAndPage, whereClause string, args []interface{}) error {
	rawDataList := make([]models.GalleryData, 0, 10)
	baseSelect := `
        SELECT g.id,g.img_url,g.kind_id,k.name
        	FROM gallery g
        	LEFT JOIN gallery_kind k ON g.kind_id = k.id
        	`

	orderBy := "ORDER BY g.id DESC"

	sqlStr := fmt.Sprintf("%s %s %s  LIMIT ? OFFSET ?",
		baseSelect, whereClause, orderBy)

	if err := r.db.Select(&rawDataList, sqlStr, args...); err != nil {
		return err
	}
	// 处理查询结果
	data.GalleryList = make([]models.GalleryData, len(rawDataList))
	for i, raw := range rawDataList {
		data.GalleryList[i] = raw
	}
	return nil
}

func (r *GalleryRepoMySQL) getCount(data *models.GalleryListAndPage, PageSize int, whereClause string, args []interface{}) error {
	baseSql := `
        SELECT COUNT(DISTINCT g.id)
        FROM gallery g  
   	`

	var totalCount int
	sqlStr := fmt.Sprintf("%s %s", baseSql, whereClause)
	if err := r.db.Get(&totalCount, sqlStr, args[:len(args)-2]...); err != nil {
		return err
	}

	data.TotalPages = (totalCount + PageSize - 1) / PageSize
	return nil
}
