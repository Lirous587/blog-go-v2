package mysql

import (
	"blog/models"
)

func CreateGalleryKind(p *models.GalleryKindParams) error {
	sqlStr := `INSERT INTO gallery_kind(name ) VALUES (:name)`
	_, err := db.NamedExec(sqlStr, p)
	return err
}

func DeleteGalleryKind(id int) error {
	sqlStr := `DELETE FROM gallery_kind WHERE id = ?`
	_, err := db.Exec(sqlStr, id)
	return err
}

func UpdateGalleryKind(p *models.GalleryKindUpdateParams) error {
	sqlStr := `UPDATE gallery_kind SET name = :name WHERE id = :id`
	_, err := db.NamedExec(sqlStr, p)
	return err
}

func GetGalleryKindList(list *models.GalleryKindList) error {
	sqlStr := `SELECT id,name FROM gallery_kind`
	err := db.Select(&list.List, sqlStr)
	return err
}
