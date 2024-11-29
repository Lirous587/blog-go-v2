package repository

import (
	"blog/models"
	"github.com/jmoiron/sqlx"
)

type EssayKindRepo interface {
	Create(data *models.EssayKindParam) error
	Update(data *models.EssayKindUpdateParam) error
	Delete(id int) error
	GetList() ([]models.EssayKindData, error)
}

type EssayKindRepoMySQL struct {
	db *sqlx.DB
}

func NewEssayKindRepoMySQL(db *sqlx.DB) *EssayKindRepoMySQL {
	return &EssayKindRepoMySQL{
		db: db,
	}
}

func (r *EssayKindRepoMySQL) Create(data *models.EssayKindParam) error {
	sqlStr := `INSERT INTO e_kind(name, icon,introduction) VALUES (:name,:icon,:introduction)`
	_, err := r.db.NamedExec(sqlStr, data)
	return err
}

func (r *EssayKindRepoMySQL) Delete(id int) error {
	sqlStr := `DELETE FROM e_kind WHERE id = ?`
	_, err := r.db.Exec(sqlStr, id)
	return err
}

func (r *EssayKindRepoMySQL) Update(data *models.EssayKindUpdateParam) error {
	sqlStr := `UPDATE e_kind SET name = :name,icon = :icon,introduction = :introduction WHERE id = :id`
	_, err := r.db.NamedExec(sqlStr, data)
	return err
}

func (r *EssayKindRepoMySQL) GetList() (list []models.EssayKindData, err error) {
	list = make([]models.EssayKindData, 0, 10)
	sqlStr := `
	SELECT k.name,k.icon,k.id,k.introduction,
	    COUNT(e.id) AS essay_count 
		FROM e_kind k
		LEFT JOIN essay e ON k.id = e.kind_id
		GROUP BY k.id
		`
	err = r.db.Select(&list, sqlStr)
	return
}
