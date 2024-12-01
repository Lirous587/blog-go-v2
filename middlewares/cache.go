package middlewares

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
)

func UpdateIndexMiddleware(indexCtrl *controller.IndexCtrl) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 先执行后续操作

		// 如果之前发生错误，不执行更新逻辑
		if len(c.Errors) > 0 {
			return
		}

		if err := indexCtrl.Update(); err != nil {
			// 处理更新错误
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}
}

func UpdateEssayDescMiddleware(essayCtrl *controller.EssayCtrl) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 先执行后续操作
		// 如果之前发生错误，不执行更新逻辑
		if len(c.Errors) > 0 {
			return
		}
		if err := essayCtrl.UpdateDescCache(); err != nil {
			// 处理更新错误
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}
}
