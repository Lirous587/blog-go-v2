package cache

import (
	"blog/models"
	"blog/repository"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var (
	rwLockForIndex sync.RWMutex
)

type IndexCache interface {
	GetData() (*models.IndexData, error)
	GetDataFromRepo() (*models.IndexData, error)
	SaveData(data *models.IndexData)
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

func (cch *IndexCacheRedis) GetDataFromRepo() (data *models.IndexData, err error) {
	// 从repo里面查询数据 然后序列化 保存到以字符串的形式保存到redis里面
	data = new(models.IndexData)
	key := getRedisKey(KeyIndex)
	data, err = cch.selectDataFromRepo()
	if err != nil {
		return
	}
	// 将IndexData结构体序列化为JSON字符串
	serializedData, marshalErr := json.Marshal(data)
	if marshalErr != nil {
		err = marshalErr
		return
	}
	// 设置过期时间，这里假设设置为600秒（10分钟），可根据实际需求调整
	expiration := time.Second * 600
	// 保存到Redis中
	setErr := cch.rdb.Set(key, string(serializedData), expiration).Err()
	if setErr != nil {
		err = setErr
		return
	}
	return
}

func (cch *IndexCacheRedis) selectDataFromRepo() (data *models.IndexData, err error) {
	ekRepo := repository.EssayKindRepo(repository.NewEssayKindRepoMySQL(repository.DB))
	kindList, err := ekRepo.GetList()
	if err != nil {
		return
	}

	elRepo := repository.EssayLabelRepo(repository.NewEssayLabelRepoMySQL(repository.DB))
	labelList, err := elRepo.GetList()
	if err != nil {
		return
	}

	repo := repository.EssayRepo(repository.NewEssayRepoMySQL(repository.DB))
	essayList, err := repo.GetRecommendList()
	if err != nil {
		return
	}

	hwRepo := repository.HeartWordsRepo(repository.NewHeartWordsRepoMySQL(repository.DB))
	heartWordsList, err := hwRepo.GetCouldTypeList()
	if err != nil {
		return
	}
	data = new(models.IndexData)
	//整合数据
	data.KindList = kindList
	data.LabelList = labelList
	data.EssayList = essayList
	data.HeartWordsList = heartWordsList
	return
}

func (cch *IndexCacheRedis) SaveData(data *models.IndexData) {
	return
}
