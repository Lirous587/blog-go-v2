package cache

import (
	"blog/dao/mysql"
	"blog/models"
	"fmt"
	"go.uber.org/zap"
	"sync"
)

var (
	rwLockForIndex       sync.RWMutex
	globalDataAboutIndex = new(models.IndexData)
)

func refreshData(data *models.IndexData) (err error) {
	var kindList = new([]models.KindData)
	if err = mysql.GetKindList(kindList); err != nil {
		return err
	}
	var labelList = new([]models.LabelData)
	if err = mysql.GetLabelList(labelList); err != nil {
		return err
	}

	var essayList = new([]models.EssayData)
	if err = mysql.GetRecommendEssayList(essayList); err != nil {
		return err
	}

	var heartWordsList = new([]models.HeartWordsData)
	if err = mysql.GetTypeHeartWords(heartWordsList); err != nil {
		return err
	}
	//整合数据
	data.KindList = *kindList
	data.LabelList = *labelList
	data.EssayList = *essayList
	data.HeartWordsList = *heartWordsList
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
