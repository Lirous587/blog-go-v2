package middlewares

import (
	"blog/cache"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sync"
)

func UpdateDataMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在请求被处理之前，不做任何事情
		// 调用下一个中间件或处理函数
		c.Next()
		var wg sync.WaitGroup
		wg.Add(2)
		errChan := make(chan error, 2)
		go func() {
			errChan <- cache.UpdateIndexData()
			wg.Done()
		}()
		go func() {
			errChan <- cache.UpdateDataAboutEssayList()
			wg.Done()
		}()
		wg.Wait()
		close(errChan)
		for err := range errChan {
			if err != nil {
				zap.L().Error("cache update happen error,err:", zap.Error(err))
			}
		}
	}
}
