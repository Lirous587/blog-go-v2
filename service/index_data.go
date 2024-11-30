package service

import (
	"blog/cache"
	"blog/models"
	"blog/repository"
	"errors"
	"github.com/go-redis/redis"
)

type IndexService interface {
	GetData() (*models.IndexData, error)
}

type IndexCacheService struct {
	cache cache.IndexCache
	repo  repository.IndexRepo
}

func NewIndexDataCacheService(cache cache.IndexCache, repo repository.IndexRepo) *IndexCacheService {
	return &IndexCacheService{
		cache: cache,
		repo:  repo,
	}
}

func (s *IndexCacheService) GetData() (data *models.IndexData, err error) {
	// 先从redis里面查
	data, err = s.cache.GetData()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return
		}
		// 没数据再从mysql里面查
		data, err = s.repo.GetData()
		if err != nil {
			return
		}
		if err = s.cache.SaveData(data); err != nil {
			return
		}
	}
	return
}
