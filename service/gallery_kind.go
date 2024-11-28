package service

import (
	"blog/models"
	"blog/repository"
)

type GalleryKindService interface {
	Create(data *models.GalleryKindParams) error
	Read(id int) (*models.GalleryKindData, error)
	Update(data *models.GalleryKindUpdateParams) error
	Delete(id int) error
	GetList() (*models.GalleryKindList, error)
}

type GalleryKindRepoService struct {
	repo repository.GalleryKindRepo
}

func NewGalleryKindRepoService(repo repository.GalleryKindRepo) *GalleryKindRepoService {
	return &GalleryKindRepoService{
		repo: repo,
	}
}

func (h *GalleryKindRepoService) Create(data *models.GalleryKindParams) error {
	return h.repo.Create(data)
}

func (h *GalleryKindRepoService) Read(id int) (data *models.GalleryKindData, err error) {
	return h.repo.Read(id)
}

func (h *GalleryKindRepoService) Update(data *models.GalleryKindUpdateParams) error {
	return h.repo.Update(data)
}

func (h *GalleryKindRepoService) Delete(id int) error {
	return h.repo.Delete(id)
}

func (h *GalleryKindRepoService) GetList() (data *models.GalleryKindList, err error) {
	return h.repo.GetList()
}
