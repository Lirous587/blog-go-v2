package mysql

import (
	"blog/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"sync"
)

type HeartWordsMysqlImpl struct {
	db *sqlx.DB
}

type HeartWordsMysql interface {
	Create(data *models.HeartWordsData) error
	Read(id int) (*models.HeartWordsData, error)
	Update(data *models.HeartWordsData) error
	Delete(id int) error
	GetList(query models.HeartWordsQuery) (*models.HeartWordsListAndPage, error)
	getList(data *models.HeartWordsListAndPage, whereClause string, args []interface{}) error
	getCount(data *models.HeartWordsListAndPage, PageSize int, whereClause string, args []interface{}) error
}

func NewHeartWordsMysql(db *sqlx.DB) HeartWordsMysql {
	return &HeartWordsMysqlImpl{
		db: db,
	}
}

func (h *HeartWordsMysqlImpl) Create(data *models.HeartWordsData) error {
	sqlStr := `INSERT INTO heart_words(content, source,img_id,if_could_type) VALUES (:content,:source,:img_id,:if_could_type)`
	_, err := h.db.NamedExec(sqlStr, data)
	return err
}

func (h *HeartWordsMysqlImpl) Read(id int) (data *models.HeartWordsData, err error) {
	sqlStr := `SELECT content,source,img_id,if_could_type FROM heart_words WHERE id = ?`
	err = h.db.Get(data, sqlStr, id)
	return data, err
}

func (h *HeartWordsMysqlImpl) Update(data *models.HeartWordsData) error {
	sqlStr := `UPDATE heart_words SET  content = :content,source = :source,img_id = :img_id, if_could_type = :if_could_type WHERE id = :id`
	_, err := h.db.NamedExec(sqlStr, data)
	return err
}

func (h *HeartWordsMysqlImpl) Delete(id int) error {
	sqlStr := `DELETE FROM heart_words WHERE id = ?`
	_, err := h.db.Exec(sqlStr, id)
	return err
}

func (h *HeartWordsMysqlImpl) GetList(query models.HeartWordsQuery) (*models.HeartWordsListAndPage, error) {
	data := &models.HeartWordsListAndPage{}
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
		if err := h.getList(data, whereClause, args); err != nil {
			errChan <- fmt.Errorf("getList failed, err: %w", err)
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := h.getCount(data, query.PageSize, whereClause, args); err != nil {
			errChan <- fmt.Errorf("getList failed, err: %w", err)
			return
		}
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

func (h *HeartWordsMysqlImpl) getList(data *models.HeartWordsListAndPage, whereClause string, args []interface{}) error {
	rawDataList := make([]models.HeartWordsData, 0, 10)
	baseSelect := `
        SELECT h.id, h.content, h.source, h.img_id, h.if_could_type, g.img_url
        FROM heart_words h
        LEFT JOIN blog.gallery g ON h.img_id = g.id
    `
	orderBy := "ORDER BY h.id DESC"

	sqlStr := fmt.Sprintf("%s %s %s LIMIT ? OFFSET ?", baseSelect, whereClause, orderBy)

	if err := h.db.Select(&rawDataList, sqlStr, args...); err != nil {
		return err
	}

	// 处理查询结果
	data.HeartWordsList = make([]models.HeartWordsData, len(rawDataList))
	for i, raw := range rawDataList {
		data.HeartWordsList[i] = raw
	}
	return nil
}

func (h *HeartWordsMysqlImpl) getCount(data *models.HeartWordsListAndPage, PageSize int, whereClause string, args []interface{}) error {
	baseSql := `
        SELECT COUNT(DISTINCT h.id)
        FROM heart_words h
    `
	var totalCount int
	sqlStr := fmt.Sprintf("%s %s", baseSql, whereClause)
	if err := h.db.Get(&totalCount, sqlStr, args[:len(args)-2]...); err != nil {
		return err
	}

	data.TotalPage = (totalCount + PageSize - 1) / PageSize
	return nil
}
