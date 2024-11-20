package cache

import (
	"blog/dao/mysql"
	"blog/dao/redis"
	"blog/models"
	"fmt"
	"go.uber.org/zap"
	"sync"
)

var (
	globalDataAboutEssayList = new([]models.EssayData)
	rwLockForEssayList       sync.RWMutex
)

func GetEssayListInit() (*[]models.EssayData, error) {
	rwLockForEssayList.Lock()
	defer rwLockForEssayList.Unlock()
	if err := mysql.GetAllEssay(globalDataAboutEssayList); err != nil {
		zap.L().Error("mysql.GetAllEssay(globalDataAboutEssayList) filed,err:", zap.Error(err))
		return nil, err
	}

	if err := redis.GetEssayKeywords(globalDataAboutEssayList); err != nil {
		zap.L().Error("redis.GetEssayKeywords(globalDataAboutEssayList) filed,err:", zap.Error(err))
		return nil, err
	}

	return globalDataAboutEssayList, nil
}

func GetAllEssayList() *[]models.EssayData {
	return globalDataAboutEssayList
}

func UpdateDataAboutEssayList() error {
	var wg sync.WaitGroup
	wg.Add(1)
	errChan := make(chan error, 1)
	go func() {
		if _, err := GetEssayListInit(); err != nil {
			errChan <- fmt.Errorf("happen err in cache UpdateDataAboutEssayList:%w", err)
		}
		wg.Done()
	}()

	wg.Wait()
	close(errChan)
	return <-errChan
}
