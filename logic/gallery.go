package logic

import (
	"blog/dao/mysql"
	"blog/models"
)

func CreateGallery(p *models.GalleryParams) error {
	return mysql.CreateGallery(p)
}

func DeleteGallery(id int) error {
	return mysql.DeleteGallery(id)
}

func UpdateGallery(p *models.GalleryUpdateParams) error {
	return mysql.UpdateGallery(p)
}

func GetGalleryList(list *models.GalleryListAndPage, query models.GalleryQuery) error {
	if err := mysql.GetGalleryList(list, query); err != nil {
		return err
	}
	return nil
}
