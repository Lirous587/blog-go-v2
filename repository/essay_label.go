package repository

import (
	"blog/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type EssayLabelRepo interface {
	Create(data *models.EssayLabelData) error
	Update(data *models.EssayLabelData) error
	Delete(id int) error
	GetList() (*[]models.EssayLabelData, error)
}

type EssayLabelRepoMySQL struct {
	db *sqlx.DB
}

func NewEssayLabelRepoMySQL(db *sqlx.DB) *EssayLabelRepoMySQL {
	return &EssayLabelRepoMySQL{
		db: db,
	}
}

func (r *EssayLabelRepoMySQL) Create(data *models.EssayLabelData) error {
	sqlStr := `INSERT INTO e_label (name,introduction) VALUES(:name,:introduction)`
	_, err := r.db.NamedExec(sqlStr, data)
	if err != nil {
		return err
	}
	return err
}

func (r *EssayLabelRepoMySQL) Update(data *models.EssayLabelData) error {
	sqlStr := `UPDATE e_label SET name = :name,introduction=:introduction WHERE id = :id`
	_, err := r.db.NamedExec(sqlStr, data)
	if err != nil {
		return fmt.Errorf("update label failed,err:%w", err)
	}
	return nil
}

func (r *EssayLabelRepoMySQL) Delete(id int) error {
	return newSqlxTx(r.db, func(tx *sqlx.Tx) error {
		if err := deleteLabelInEssayLabel(tx, id); err != nil {
			return fmt.Errorf("deleteLabelFromEssayLabel failed,err:%w", err)
		}
		if err := deleteLabelInLabel(tx, id); err != nil {
			return fmt.Errorf("deleteLabelFromLabel failed,err:%w", err)
		}
		return nil
	})
}

func deleteLabelInLabel(tx *sqlx.Tx, id int) error {
	// 删除label
	sqlStr := `DELETE FROM e_label WHERE id = ?`
	_, err := tx.Exec(sqlStr, id)
	return err
}

func deleteLabelInEssayLabel(tx *sqlx.Tx, id int) error {
	sqlStr := `DELETE FROM eid_lid WHERE label_id = ?`
	_, err := tx.Exec(sqlStr, id)
	return err
}

func (r *EssayLabelRepoMySQL) GetList() (data *[]models.EssayLabelData, err error) {
	data = new([]models.EssayLabelData)
	*data = make([]models.EssayLabelData, 0, 10)
	sqlStr := `
		SELECT l.name,l.id,l.introduction,
		    COUNT(el.essay_id) AS essay_count 
			FROM e_label l
			LEFT JOIN eid_lid el on l.id = el.label_id
			GROUP BY l.id
			`
	err = r.db.Select(data, sqlStr)
	return
}
