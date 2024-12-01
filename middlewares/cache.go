package middlewares

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
)

func UpdateDataMiddleware(indexCtrl *controller.IndexCtrl) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 先执行后续操作
		if err := indexCtrl.Update(); err != nil {
			// 处理更新错误
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}
}
