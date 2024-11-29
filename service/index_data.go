package service

import (
	"blog/cache"
	"blog/models"
)

type IndexService interface {
	GetData() error
}

type IndexCacheService struct {
}

func GetIndexData(data **models.IndexData) {
	// 从缓存中拿到数据
	cache.GetIndexData(data)
}
