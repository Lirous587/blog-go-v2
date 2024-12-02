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

func (s *GalleryKindRepoService) Create(data *models.GalleryKindParams) error {
	return s.repo.Create(data)
}

func (s *GalleryKindRepoService) Read(id int) (data *models.GalleryKindData, err error) {
	return s.repo.Read(id)
}

func (s *GalleryKindRepoService) Update(data *models.GalleryKindUpdateParams) error {
	return s.repo.Update(data)
}

func (s *GalleryKindRepoService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *GalleryKindRepoService) GetList() (data *models.GalleryKindList, err error) {
	return s.repo.GetList()
}
