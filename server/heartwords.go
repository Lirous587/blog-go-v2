package server

import (
	"blog/dao/mysql"
	"blog/models"
)

type HeartWordsServer interface {
	Create(data *models.HeartWordsData) error
	Read(id int) (*models.HeartWordsData, error)
	Update(data *models.HeartWordsData) error
	Delete(id int) error
	GetList(query models.HeartWordsQuery) (*models.HeartWordsListAndPage, error)
}

type HeartWordsImpl struct {
	repo mysql.HeartWordsMysql
}

func NewHeartWordsServer(repo mysql.HeartWordsMysql) HeartWordsServer {
	return &HeartWordsImpl{
		repo: repo,
	}
}

func (h *HeartWordsImpl) Create(data *models.HeartWordsData) error {
	return h.repo.Create(data)
}

func (h *HeartWordsImpl) Read(id int) (data *models.HeartWordsData, err error) {
	return h.repo.Read(id)
}

func (h *HeartWordsImpl) Update(data *models.HeartWordsData) error {
	return h.repo.Update(data)
}

func (h *HeartWordsImpl) Delete(id int) error {
	return h.repo.Delete(id)
}

func (h *HeartWordsImpl) GetList(query models.HeartWordsQuery) (data *models.HeartWordsListAndPage, err error) {
	return h.repo.GetList(query)
}
