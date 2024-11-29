package cache

import (
	"blog/models"
	"blog/repository"
	"fmt"
	"go.uber.org/zap"
	"sync"
)

var (
	rwLockForIndex       sync.RWMutex
	globalDataAboutIndex = new(models.IndexData)
)

func refreshData(data *models.IndexData) error {
	ekRepo := repository.EssayKindRepo(repository.NewEssayKindRepoMySQL(repository.DB))
	kindList, err := ekRepo.GetList()
	if err != nil {
		return err
	}

	elRepo := repository.EssayLabelRepo(repository.NewEssayLabelRepoMySQL(repository.DB))
	labelList, err := elRepo.GetList()
	if err != nil {
		return err
	}

	repo := repository.EssayRepo(repository.NewEssayRepoMySQL(repository.DB))
	essayList, err := repo.GetRecommendList()
	if err != nil {
		return err
	}

	hwRepo := repository.HeartWordsRepo(repository.NewHeartWordsRepoMySQL(repository.DB))
	heartWordsList, err := hwRepo.GetCouldTypeList()
	if err != nil {
		return err
	}
	//整合数据
	data.KindList = kindList
	data.LabelList = labelList
	data.EssayList = essayList
	data.HeartWordsList = heartWordsList
	return nil
}

func InitIndexData() error {
	rwLockForIndex.Lock()
	defer rwLockForIndex.Unlock()
	if err := refreshData(globalDataAboutIndex); err != nil {
		zap.L().Error("help.ResponseDataAboutIndex(globalDataAboutIndex) failed,err:", zap.Error(err))
		return err
	}
	return nil
}

func GetIndexData(data **models.IndexData) {
	*data = globalDataAboutIndex
}

func UpdateIndexData() error {
	var wg sync.WaitGroup
	wg.Add(1)
	errChan := make(chan error, 1)
	go func() {
		defer wg.Done()
		if err := InitIndexData(); err != nil {
			errChan <- fmt.Errorf("happen err in cache UpdateDataAboutIndex:%w", err)
		}
	}()

	wg.Wait()
	close(errChan)
	return <-errChan
}
