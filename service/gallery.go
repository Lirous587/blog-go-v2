package service

import (
	"blog/models"
	"blog/repository"
)

type GalleryService interface {
	Create(data *models.GalleryParams) error
	Read(id int) (*models.GalleryData, error)
	Update(data *models.GalleryUpdateParams) error
	Delete(id int) error
	GetList(query *models.GalleryQuery) (*models.GalleryListAndPage, error)
}

type GalleryRepoService struct {
	repo repository.GalleryRepo
}

func NewGalleryRepoService(repo repository.GalleryRepo) *GalleryRepoService {
	return &GalleryRepoService{
		repo: repo,
	}
}

func (s *GalleryRepoService) Create(data *models.GalleryParams) error {
	return s.repo.Create(data)
}

func (s *GalleryRepoService) Read(id int) (data *models.GalleryData, err error) {
	return s.repo.Read(id)
}

func (s *GalleryRepoService) Update(data *models.GalleryUpdateParams) error {
	return s.repo.Update(data)
}

func (s *GalleryRepoService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *GalleryRepoService) GetList(query *models.GalleryQuery) (data *models.GalleryListAndPage, err error) {
	return s.repo.GetList(query)
}
