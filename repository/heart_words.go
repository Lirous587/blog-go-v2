package repository

import (
	"blog/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"sync"
)

type HeartWordsRepo interface {
	Create(data *models.HeartWordsParam) error
	Read(id int) (*models.HeartWordsData, error)
	Update(data *models.HeartWordsUpdateParam) error
	Delete(id int) error
	GetList(query *models.HeartWordsQuery) (*models.HeartWordsListAndPage, error)
	GetCouldTypeList() ([]models.HeartWordsData, error)
}

type HeartWordsRepoMySQL struct {
	db *sqlx.DB
}

func NewHeartWordsRepoMySQL(db *sqlx.DB) *HeartWordsRepoMySQL {
	return &HeartWordsRepoMySQL{
		db: db,
	}
}

func (r *HeartWordsRepoMySQL) Create(data *models.HeartWordsParam) error {
	sqlStr := `INSERT INTO heart_words(content, source,img_id,if_could_type) VALUES (:content,:source,:img_id,:if_could_type)`
	_, err := r.db.NamedExec(sqlStr, data)
	return err
}

func (r *HeartWordsRepoMySQL) Read(id int) (data *models.HeartWordsData, err error) {
	sqlStr := `SELECT content,source,img_id,if_could_type FROM heart_words WHERE id = ?`
	err = r.db.Get(data, sqlStr, id)
	return data, err
}

func (r *HeartWordsRepoMySQL) Update(data *models.HeartWordsUpdateParam) error {
	sqlStr := `UPDATE heart_words SET  content = :content,source = :source,img_id = :img_id, if_could_type = :if_could_type WHERE id = :id`
	_, err := r.db.NamedExec(sqlStr, data)
	return err
}

func (r *HeartWordsRepoMySQL) Delete(id int) error {
	sqlStr := `DELETE FROM heart_words WHERE id = ?`
	_, err := r.db.Exec(sqlStr, id)
	return err
}

func (r *HeartWordsRepoMySQL) GetList(query *models.HeartWordsQuery) (*models.HeartWordsListAndPage, error) {
	data := new(models.HeartWordsListAndPage)
	var wg sync.WaitGroup
	taskCount := 2
	var errChan = make(chan error, taskCount)
	wg.Add(taskCount)

	// 计算偏移量
	offset := (query.Page - 1) * query.PageSize
	args := make([]interface{}, 0)

	where := make([]string, 0)
	if query.IfCouldType {
		where = append(where, "h.if_could_type = 1")
	}
	// 为后续查询筛选条件做准备
	whereClause := ""
	if len(where) > 0 {
		whereClause = "WHERE " + strings.Join(where, " AND ")
	}

	args = append(args, query.PageSize, offset)

	go func() {
		defer wg.Done()
		list, err := r.getList(whereClause, args)
		if err != nil {
			errChan <- fmt.Errorf("getList failed, err: %w", err)
			return
		}
		data.HeartWordsList = *list
	}()

	go func() {
		defer wg.Done()
		totalCount, err := r.getCount(whereClause, args)
		if err != nil {
			errChan <- fmt.Errorf("getList failed, err: %w", err)
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

func (r *HeartWordsRepoMySQL) GetCouldTypeList() (data []models.HeartWordsData, err error) {
	data = make([]models.HeartWordsData, 0, 10)
	sqlStr :=
		`SELECT h.id, h.content, h.source, h.img_id, h.if_could_type, g.img_url FROM heart_words h 
			LEFT JOIN blog.gallery g ON h.img_id = g.id
			WHERE h.if_could_type = 1`
	err = r.db.Select(data, sqlStr)
	return
}

func (r *HeartWordsRepoMySQL) getList(whereClause string, args []interface{}) (list *[]models.HeartWordsData, err error) {
	list = new([]models.HeartWordsData)
	*list = make([]models.HeartWordsData, 0, 10)
	baseSelect := `
        SELECT h.id, h.content, h.source, h.img_id, h.if_could_type, g.img_url
        FROM heart_words h
        LEFT JOIN blog.gallery g ON h.img_id = g.id
    `
	orderBy := "ORDER BY h.id DESC"

	sqlStr := fmt.Sprintf("%s %s %s LIMIT ? OFFSET ?", baseSelect, whereClause, orderBy)

	if err = r.db.Select(list, sqlStr, args...); err != nil {
		return
	}
	return
}

func (r *HeartWordsRepoMySQL) getCount(whereClause string, args []interface{}) (totalCount int, err error) {
	baseSql := `
        SELECT COUNT(DISTINCT h.id)
        FROM heart_words h
    `
	sqlStr := fmt.Sprintf("%s %s", baseSql, whereClause)
	if err = r.db.Get(&totalCount, sqlStr, args[:len(args)-2]...); err != nil {
		return
	}
	return
}
