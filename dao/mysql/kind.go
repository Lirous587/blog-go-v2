package mysql

import "blog/models"

func GetKindList(data *[]models.KindData) error {
	sqlStr := `
	SELECT k.name,k.icon,k.id,k.introduction,
	    COUNT(e.id) AS essay_count 
		FROM kind k
		LEFT JOIN essay e ON k.id = e.kind_id
		GROUP BY k.id
		`
	return db.Select(data, sqlStr)
}

func CreateKind(p *models.KindParams) error {
	sqlStr := `INSERT INTO kind(name, icon,introduction) VALUES (:name,:icon,:introduction)`
	_, err := db.NamedExec(sqlStr, p)
	return err
}

func DeleteKind(id int) error {
	sqlStr := `DELETE FROM kind WHERE id = ?`
	_, err := db.Exec(sqlStr, id)
	return err
}

func UpdateKind(p *models.KindUpdateParams) error {
	sqlStr := `UPDATE kind SET name = :name,icon = :icon,introduction = :introduction WHERE id = :id`
	_, err := db.NamedExec(sqlStr, p)
	return err
}
