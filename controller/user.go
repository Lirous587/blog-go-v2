package controller

import (
	"blog/models"
	"blog/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

const (
	signupSuccess        = "注册成功"
	userIDInvalid        = "无法获取该用户id"
	updateUserMsgSuccess = "修改个人信息成功"
)

// SignupHandler 注册
func SignupHandler(c *gin.Context) {
	//1.获取参数和参数校验
	var p = new(models.UserParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.业务处理
	if err := service.Signup(p); err != nil {
		zap.L().Error("logic.Signup(p) failed", zap.Error(err))
		ResponseError(c, CodeUserExist)
		return
	}
	//3.返回响应
	ResponseSuccess(c, signupSuccess)
}

// LoginHandler 登录
func LoginHandler(c *gin.Context) {
	//1.获取参数并检验
	var p = new(models.User)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.业务处理
	if err := service.Login(p); err != nil {
		zap.L().Error("logic.Login() failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, gin.H{
		"token": p.Token,
	})
}

// LogoutHandler 退出登录
func LogoutHandler(c *gin.Context) {
	//1.参数验证 --> 得到相应的token
	authHeader := c.Request.Header.Get("Authorization")
	parts := strings.SplitN(authHeader, " ", 2)
	//得到token
	token := parts[1]

	//2.业务处理 --> 将该token储存在数据库中
	if err := service.Logout(token); err != nil {
		zap.L().Error("logic.Logout(token) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, CodeSuccess)
}

// UpdateUserMsgHandler 修改用户信息
func UpdateUserMsgHandler(c *gin.Context) {
	//1.参数校验
	var p = new(models.UserParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//获取id
	id, err := getUserId(c)
	if err != nil {
		zap.L().Error("getUserId(c) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	//2.业务处理
	if err = service.UpdateUserMsg(p, id); err != nil {
		zap.L().Error("logic.UpdateUserMsg(user, id) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	//3.返回响应
	ResponseSuccess(c, updateUserMsgSuccess)
}

const CtxUserIDKey = "UserID"

func getUserId(c *gin.Context) (id int64, err error) {
	uid, exist := c.Get(CtxUserIDKey)
	if !exist {
		return 0, err
	}
	var ok bool
	id, ok = uid.(int64)
	if !ok {
		return 0, err
	}
	return
}
