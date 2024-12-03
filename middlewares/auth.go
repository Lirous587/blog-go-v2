package middlewares

import (
	"blog/cache"
	"blog/pkg/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authFailedMsg = "需要登录"
)

func NewUserAuthMiddleware(cch cache.UserCache) func(c *gin.Context) {
	return func(c *gin.Context) {
		cookies := c.Request.CookiesNamed("user_token")
		fmt.Println(cookies)
		//token := cookies[0].Value
		//if token == "" {
		//	c.JSON(http.StatusInternalServerError, gin.H{
		//		"msg": authFailedMsg,
		//	})
		//	c.Abort()
		//	return
		//}
		//
		//uData, err := cch.ParseToken(token)
		//if err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{
		//		"msg": authFailedMsg,
		//	})
		//	c.Abort()
		//	return
		//}
		//c.Set(controller.CtxUserIDKey, uData)
		c.Next()
	}
}

func NewManagerAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer xxxxxxx.xxx.xxx  / X-TOKEN: xxx.xxx.xx
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			//controller.ResponseError(c, controller.CodeNeedLogin)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": authFailedMsg,
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": authFailedMsg,
			})
			c.Abort()
			return
		}

		_, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": authFailedMsg,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
