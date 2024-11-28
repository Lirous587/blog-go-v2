package repository

import (
	"blog/models"
	"github.com/jmoiron/sqlx"
)

type GalleryKindRepoMySQL struct {
	db *sqlx.DB
}

type GalleryKindRepo interface {
	Create(data *models.GalleryKindParams) error
	Read(id int) (*models.GalleryKindData, error)
	Update(data *models.GalleryKindUpdateParams) error
	Delete(id int) error
	GetList() (*models.GalleryKindList, error)
}

func NewGalleryKindRepoMySQL(db *sqlx.DB) *GalleryKindRepoMySQL {
	return &GalleryKindRepoMySQL{
		db: db,
	}
}

func (r *GalleryKindRepoMySQL) Create(data *models.GalleryKindParams) error {
	sqlStr := `INSERT INTO gallery_kind(name) VALUES (:name)`
	_, err := r.db.NamedExec(sqlStr, data)
	return err
}

func (r *GalleryKindRepoMySQL) Read(id int) (data *models.GalleryKindData, err error) {
	sqlStr := `SELECT id,name  FROM gallery_kind  where id = ?`
	err = r.db.Select(data, sqlStr, id)
	return
}

func (r *GalleryKindRepoMySQL) Update(data *models.GalleryKindUpdateParams) error {
	sqlStr := `UPDATE gallery_kind SET name = :name WHERE id = :id`
	_, err := r.db.NamedExec(sqlStr, data)
	return err
}

func (r *GalleryKindRepoMySQL) Delete(id int) error {
	sqlStr := `DELETE FROM gallery_kind WHERE id = ?`
	_, err := r.db.Exec(sqlStr, id)
	return err
}

func (r *GalleryKindRepoMySQL) GetList() (data *models.GalleryKindList, err error) {
	data = new(models.GalleryKindList)
	sqlStr := `SELECT id,name FROM gallery_kind`
	err = r.db.Select(&data.List, sqlStr)
	return data, err
}
