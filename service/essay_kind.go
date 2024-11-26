package service

import (
	"blog/models"
	"blog/repository"
)

type EssayKindService interface {
	Create(data *models.EssayKindData) error
	Delete(id int) error
	Update(data *models.EssayKindData) error
}

type EssayKindRepoService struct {
	repo repository.EssayKindRepo
}

func NewEssayKindRepoService(repo repository.EssayKindRepo) EssayKindRepoService {
	return EssayKindRepoService{
		repo: repo,
	}
}

func (s EssayKindRepoService) Create(data *models.EssayKindData) error {
	return s.repo.Create(data)
}

func (s EssayKindRepoService) Update(data *models.EssayKindData) error {
	return s.repo.Update(data)
}

func (s EssayKindRepoService) Delete(id int) error {
	return s.repo.Delete(id)
}
