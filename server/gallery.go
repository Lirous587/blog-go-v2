package server

import (
	"blog/dao/mysql"
	"blog/models"
)

type GalleryServer interface {
	Create(data *models.GalleryData) error
	Read(id int) (*models.GalleryData, error)
	Update(data *models.GalleryData) error
	Delete(id int) error
	GetList(query models.GalleryQuery) (*models.GalleryListAndPage, error)
}

type GalleryImpl struct {
	repo mysql.GalleryMysql
}

func NewGalleryServer(repo mysql.GalleryMysql) GalleryServer {
	return &GalleryImpl{
		repo: repo,
	}
}

func (h *GalleryImpl) Create(data *models.GalleryData) error {
	return h.repo.Create(data)
}

func (h *GalleryImpl) Read(id int) (data *models.GalleryData, err error) {
	return h.repo.Read(id)
}

func (h *GalleryImpl) Update(data *models.GalleryData) error {
	return h.repo.Update(data)
}

func (h *GalleryImpl) Delete(id int) error {
	return h.repo.Delete(id)
}

func (h *GalleryImpl) GetList(query models.GalleryQuery) (data *models.GalleryListAndPage, err error) {
	return h.repo.GetList(query)
}
