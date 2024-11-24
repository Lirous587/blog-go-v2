package server

import (
	"blog/models"
	"blog/repository"
)

type HeartWordsServer interface {
	Create(data *models.HeartWordsData) error
	Read(id int) (*models.HeartWordsData, error)
	Update(data *models.HeartWordsData) error
	Delete(id int) error
	GetList(query models.HeartWordsQuery) (*models.HeartWordsListAndPage, error)
}

type RepoHeartWordsService struct {
	repo repository.HeartWordsRepo
}

func NewRepoHeartWordsService(repo repository.HeartWordsRepo) *RepoHeartWordsService {
	return &RepoHeartWordsService{
		repo: repo,
	}
}

func (h *RepoHeartWordsService) Create(data *models.HeartWordsData) error {
	return h.repo.Create(data)
}

func (h *RepoHeartWordsService) Read(id int) (data *models.HeartWordsData, err error) {
	return h.repo.Read(id)
}

func (h *RepoHeartWordsService) Update(data *models.HeartWordsData) error {
	return h.repo.Update(data)
}

func (h *RepoHeartWordsService) Delete(id int) error {
	return h.repo.Delete(id)
}

func (h *RepoHeartWordsService) GetList(query models.HeartWordsQuery) (data *models.HeartWordsListAndPage, err error) {
	return h.repo.GetList(query)
}
