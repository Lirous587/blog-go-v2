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

func (h *HeartWordsRepoService) Create(data *models.HeartWordsParam) error {
	return h.repo.Create(data)
}

func (h *HeartWordsRepoService) Read(id int) (data *models.HeartWordsData, err error) {
	return h.repo.Read(id)
}

func (h *HeartWordsRepoService) Update(data *models.HeartWordsUpdateParam) error {
	return h.repo.Update(data)
}

func (h *HeartWordsRepoService) Delete(id int) error {
	return h.repo.Delete(id)
}

func (h *HeartWordsRepoService) GetList(query *models.HeartWordsQuery) (data *models.HeartWordsListAndPage, err error) {
	return h.repo.GetList(query)
}
