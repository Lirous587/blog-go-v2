package service

import (
	"blog/models"
	"blog/repository"
)

type EssayLabelService interface {
	Create(data *models.EssayLabelData) error
	Update(data *models.EssayLabelData) error
	Delete(id int) error
}

type EssayLabelRepoService struct {
	repo repository.EssayLabelRepo
}

func NewEssayLabelRepoService(repo repository.EssayLabelRepo) *EssayLabelRepoService {
	return &EssayLabelRepoService{
		repo: repo,
	}
}

func (s *EssayLabelRepoService) Create(data *models.EssayLabelData) (err error) {
	return s.repo.Create(data)
}

func (s *EssayLabelRepoService) Delete(id int) (err error) {
	return s.repo.Delete(id)
}

func (s *EssayLabelRepoService) Update(data *models.EssayLabelData) (err error) {
	return s.repo.Update(data)
}