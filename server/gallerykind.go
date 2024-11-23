package server

import (
	"blog/dao/mysql"
	"blog/models"
)

type GalleryKindServer interface {
	Create(data *models.GalleryKindData) error
	Read(id int) (*models.GalleryKindData, error)
	Update(data *models.GalleryKindData) error
	Delete(id int) error
	GetList() (*models.GalleryKindList, error)
}

type GalleryKindImpl struct {
	repo mysql.GalleryKindMysql
}

func NewGalleryKindServer(repo mysql.GalleryKindMysql) GalleryKindServer {
	return &GalleryKindImpl{
		repo: repo,
	}
}

func (h *GalleryKindImpl) Create(data *models.GalleryKindData) error {
	return h.repo.Create(data)
}

func (h *GalleryKindImpl) Read(id int) (data *models.GalleryKindData, err error) {
	return h.repo.Read(id)
}

func (h *GalleryKindImpl) Update(data *models.GalleryKindData) error {
	return h.repo.Update(data)
}

func (h *GalleryKindImpl) Delete(id int) error {
	return h.repo.Delete(id)
}

func (h *GalleryKindImpl) GetList() (data *models.GalleryKindList, err error) {
	return h.repo.GetList()
}
