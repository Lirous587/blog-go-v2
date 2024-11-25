package server

import (
	"blog/models"
	"blog/repository"
)

type GalleryServer interface {
	Create(data *models.GalleryData) error
	Read(id int) (*models.GalleryData, error)
	Update(data *models.GalleryData) error
	Delete(id int) error
	GetList(query *models.GalleryQuery) (*models.GalleryListAndPage, error)
}

type RepoGalleryService struct {
	repo repository.GalleryRepo
}

func NewRepoGalleryService(repo repository.GalleryRepo) *RepoGalleryService {
	return &RepoGalleryService{
		repo: repo,
	}
}

func (s *RepoGalleryService) Create(data *models.GalleryData) error {
	return s.repo.Create(data)
}

func (s *RepoGalleryService) Read(id int) (data *models.GalleryData, err error) {
	return s.repo.Read(id)
}

func (s *RepoGalleryService) Update(data *models.GalleryData) error {
	return s.repo.Update(data)
}

func (s *RepoGalleryService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *RepoGalleryService) GetList(query *models.GalleryQuery) (data *models.GalleryListAndPage, err error) {
	return s.repo.GetList(query)
}
