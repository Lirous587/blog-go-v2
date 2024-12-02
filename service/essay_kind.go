package service

import (
	"blog/models"
	"blog/repository"
)

type EssayKindService interface {
	Create(data *models.EssayKindParam) error
	Delete(id int) error
	Update(data *models.EssayKindUpdateParam) error
}

type EssayKindRepoService struct {
	repo repository.EssayKindRepo
}

func NewEssayKindRepoService(repo repository.EssayKindRepo) *EssayKindRepoService {
	return &EssayKindRepoService{
		repo: repo,
	}
}

func (s *EssayKindRepoService) Create(data *models.EssayKindParam) error {
	return s.repo.Create(data)
}

func (s *EssayKindRepoService) Update(data *models.EssayKindUpdateParam) error {
	return s.repo.Update(data)
}

func (s *EssayKindRepoService) Delete(id int) error {
	return s.repo.Delete(id)
}
