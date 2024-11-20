package logic

import (
	"blog/dao/mysql"
	"blog/models"
)

func CreateGalleryKind(p *models.GalleryKindParams) error {
	return mysql.CreateGalleryKind(p)
}

func DeleteGalleryKind(id int) error {
	return mysql.DeleteGalleryKind(id)
}

func UpdateGalleryKind(p *models.GalleryKindUpdateParams) error {
	return mysql.UpdateGalleryKind(p)
}

func GetGalleryKindList(list *models.GalleryKindList) error {
	if err := mysql.GetGalleryKindList(list); err != nil {
		return err
	}
	return nil
}
