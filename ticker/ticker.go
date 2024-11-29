package ticker

import (
	"go.uber.org/zap"
	"sync"
	"time"
)

const (
	cleanLowFrequentKeyword = time.Hour * 24 * 30
	taskCount               = 2
)

var (
	wg    sync.WaitGroup
	errCh = make(chan error, taskCount)
)

func Init() {
	wg.Add(taskCount + 1) // 为所有任务和错误处理 goroutine 加 1

	taskList := []func() error{cleanLowFrequentKeywordTask}

	// 启动错误处理goroutine
	go func() {
		defer wg.Done()
		handleErrors()
	}()

	// 启动任务
	for _, task := range taskList {
		go func(t func() error) {
			defer wg.Done()
			if err := runTask(t); err != nil {
				errCh <- err
			}
		}(task)
	}
	// 等待所有任务完成
	go func() {
		wg.Wait()
		close(errCh) // 所有任务完成后关闭错误通道
	}()
}

func runTask(task func() error) error {
	return task()
}

func handleErrors() {
	for err := range errCh {
		// 使用你的日志库记录错误，这里用zap作为示例
		zap.L().Error("Task happen error", zap.Error(err))
	}
}

func cleanLowFrequentKeywordTask() error {
	ticker := time.NewTicker(cleanLowFrequentKeyword)
	defer ticker.Stop()
	//for range ticker.C {
	//	err := redis.CleanLowerKeywordsZsetEveryMonth()
	//	if err != nil {
	//		return err
	//	}
	//
	//}
	return nil
}
