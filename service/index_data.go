package service

import (
	"blog/cache"
	"blog/models"
)

type IndexService interface {
	GetData() (*models.IndexData, error)
}

type IndexCacheService struct {
	cache cache.IndexCache
}

func NewIndexDataCacheService(cache cache.IndexCache) *IndexCacheService {
	return &IndexCacheService{
		cache: cache,
	}
}

func (c *IndexCacheService) GetData() (*models.IndexData, error) {
	// 先从redis里面查
	data, err := c.cache.GetData()
	if err != nil {
		// 没数据再从mysql里面查
		data, err = c.cache.GetDataFromRepo()
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}
