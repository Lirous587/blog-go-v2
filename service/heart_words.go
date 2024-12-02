package service

import (
	"blog/models"
	"blog/repository"
)

type HeartWordsService interface {
	Create(data *models.HeartWordsParam) error
	Read(id int) (*models.HeartWordsData, error)
	Update(data *models.HeartWordsUpdateParam) error
	Delete(id int) error
	GetList(query *models.HeartWordsQuery) (*models.HeartWordsListAndPage, error)
}

type HeartWordsRepoService struct {
	repo repository.HeartWordsRepo
}

func NewHeartWordsRepoService(repo repository.HeartWordsRepo) *HeartWordsRepoService {
	return &HeartWordsRepoService{
		repo: repo,
	}
}

func (s *HeartWordsRepoService) Create(data *models.HeartWordsParam) error {
	return s.repo.Create(data)
}

func (s *HeartWordsRepoService) Read(id int) (data *models.HeartWordsData, err error) {
	return s.repo.Read(id)
}

func (s *HeartWordsRepoService) Update(data *models.HeartWordsUpdateParam) error {
	return s.repo.Update(data)
}

func (s *HeartWordsRepoService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *HeartWordsRepoService) GetList(query *models.HeartWordsQuery) (data *models.HeartWordsListAndPage, err error) {
	return s.repo.GetList(query)
}
