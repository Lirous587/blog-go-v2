package server

import (
	"blog/models"
	"blog/repository"
)

type GalleryKindServer interface {
	Create(data *models.GalleryKindData) error
	Read(id int) (*models.GalleryKindData, error)
	Update(data *models.GalleryKindData) error
	Delete(id int) error
	GetList() (*models.GalleryKindList, error)
}

type RepoGalleryKindServer struct {
	repo repository.GalleryKindRepo
}

func NewRepoGalleryKindServer(repo repository.GalleryKindRepo) *RepoGalleryKindServer {
	return &RepoGalleryKindServer{
		repo: repo,
	}
}

func (h *RepoGalleryKindServer) Create(data *models.GalleryKindData) error {
	return h.repo.Create(data)
}

func (h *RepoGalleryKindServer) Read(id int) (data *models.GalleryKindData, err error) {
	return h.repo.Read(id)
}

func (h *RepoGalleryKindServer) Update(data *models.GalleryKindData) error {
	return h.repo.Update(data)
}

func (h *RepoGalleryKindServer) Delete(id int) error {
	return h.repo.Delete(id)
}

func (h *RepoGalleryKindServer) GetList() (data *models.GalleryKindList, err error) {
	return h.repo.GetList()
}
