package cache

import (
	"blog/models"
	"encoding/json"
	"github.com/go-redis/redis"
	"strconv"
)

type EssayCache interface {
	GetDesc() ([]models.EssayDesc, error)
	SaveDesc([]models.EssayDesc) error
}

type EssayCacheRedis struct {
	rdb *redis.Client
}

func NewEssayCacheRedis(rdb *redis.Client) *EssayCacheRedis {
	return &EssayCacheRedis{
		rdb: rdb,
	}
}

func (cch *EssayCacheRedis) GetDesc() ([]models.EssayDesc, error) {
	key := getRedisKey(KeyEssayKeyword)

	eMap, err := cch.rdb.HGetAll(key).Result()
	if err != nil {
		return nil, err
	}
	if len(eMap) == 0 {
		return nil, redis.Nil
	}

	list := make([]models.EssayDesc, 0, len(eMap))

	for _, v := range eMap {
		var item models.EssayDesc
		err = json.Unmarshal([]byte(v), &item)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (cch *EssayCacheRedis) SaveDesc(data []models.EssayDesc) error {
	key := getRedisKey(KeyEssayKeyword)
	pipe := cch.rdb.Pipeline()

	for _, item := range data {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return err
		}
		pipe.HSet(key, strconv.Itoa(item.ID), jsonData)
	}
	_, err := pipe.Exec()
	return err
}
