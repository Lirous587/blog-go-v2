package cache

import (
	"blog/models"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
)

type IndexCache interface {
	GetData() (*models.IndexData, error)
	SaveData(data *models.IndexData) error
}

type IndexCacheRedis struct {
	rdb *redis.Client
}

func NewIndexCacheRedis(rdb *redis.Client) *IndexCacheRedis {
	return &IndexCacheRedis{
		rdb: rdb,
	}
}

func (cch *IndexCacheRedis) GetData() (data *models.IndexData, err error) {
	data = new(models.IndexData)
	key := getRedisKey(KeyIndex)
	result, err := cch.rdb.Get(key).Result()
	if errors.Is(err, redis.Nil) {
		return
	}
	// 反序列化操作
	err = json.Unmarshal([]byte(result), data)
	if err != nil {
		return
	}
	return
}

func (cch *IndexCacheRedis) SaveData(data *models.IndexData) error {
	key := getRedisKey(KeyIndex)
	// 将IndexData结构体序列化为JSON字符串
	serializedData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = cch.rdb.Set(key, string(serializedData), 0).Err()
	return err
}
