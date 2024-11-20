package middlewares

import (
	"blog/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SaveUserIp() func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := logic.SaveUserIp(c.ClientIP()); err != nil {
			zap.L().Error("logic.SaveUserIp(c.ClientIP())", zap.Error(err))
		}
		c.Next() // 后续的处理请求的函数中 可以用过c.Get(CtxUserIDKey) 来获取当前请求的用户信息
	}
}
