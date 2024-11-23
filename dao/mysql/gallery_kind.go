package mysql

import (
	"blog/models"
	"github.com/jmoiron/sqlx"
)

type GalleryKindMysqlImpl struct {
	db *sqlx.DB
}

func NewGalleryKindMysql(db *sqlx.DB) GalleryKindMysql {
	return &GalleryKindMysqlImpl{
		db: db,
	}
}

type GalleryKindMysql interface {
	Create(data *models.GalleryKindData) error
	Read(id int) (*models.GalleryKindData, error)
	Update(data *models.GalleryKindData) error
	Delete(id int) error
	GetList() (*models.GalleryKindList, error)
}

func (g *GalleryKindMysqlImpl) Create(data *models.GalleryKindData) error {
	sqlStr := `INSERT INTO gallery_kind(name ) VALUES (:name)`
	_, err := g.db.NamedExec(sqlStr, data)
	return err
}

func (g *GalleryKindMysqlImpl) Read(id int) (data *models.GalleryKindData, err error) {
	sqlStr := `SELECT id,name  FROM gallery_kind  where id = ?`
	err = g.db.Select(data, sqlStr, id)
	return
}

func (g *GalleryKindMysqlImpl) Update(data *models.GalleryKindData) error {
	sqlStr := `UPDATE gallery_kind SET name = :name WHERE id = :id`
	_, err := g.db.NamedExec(sqlStr, data)
	return err
}

func (g *GalleryKindMysqlImpl) Delete(id int) error {
	sqlStr := `DELETE FROM gallery_kind WHERE id = ?`
	_, err := db.Exec(sqlStr, id)
	return err
}

func (g *GalleryKindMysqlImpl) GetList() (data *models.GalleryKindList, err error) {
	data = new(models.GalleryKindList)
	sqlStr := `SELECT id,name FROM gallery_kind`
	err = g.db.Select(&data.List, sqlStr)
	return data, err
}
