package mysql

import (
	"blog/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func GetLabelList(data *[]models.LabelData) error {
	sqlStr := `
		SELECT l.name,l.id,l.introduction,
		    COUNT(el.essay_id) AS essay_count 
			FROM label l
			LEFT JOIN essay_label el on l.id = el.label_id
			GROUP BY l.id
			`
	return db.Select(data, sqlStr)
}

func CreateLabel(p *models.LabelParams) error {
	sqlStr := `INSERT INTO label (name,introduction) VALUES(:name,:introduction)`
	_, err := db.NamedExec(sqlStr, p)
	if err != nil {
		return err
	}
	return err
}

func DeleteLabel(id int) error {
	return withTx(func(tx *sqlx.Tx) error {
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
	sqlStr := `DELETE FROM label WHERE id = ?`
	_, err := tx.Exec(sqlStr, id)
	return err
}

func deleteLabelInEssayLabel(tx *sqlx.Tx, id int) error {
	sqlStr := `DELETE FROM essay_label WHERE label_id = ?`
	_, err := tx.Exec(sqlStr, id)
	return err
}

func UpdateLabel(p *models.LabelUpdateParams) error {
	sqlStr := `UPDATE label SET name = :name,introduction=:introduction WHERE id = :id`
	_, err := db.NamedExec(sqlStr, p)
	if err != nil {
		return fmt.Errorf("update label failed,err:%w", err)
	}
	return nil
}
