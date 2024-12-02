package controller

import (
	"blog/models"
	"blog/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	signupSuccess        = "注册成功"
	userIDInvalid        = "无法获取该用户id"
	updateUserMsgSuccess = "修改个人信息成功"
)

type UserCtrl struct {
	service service.UsrService
}

func NewUserCtrl(service service.UsrService) *UserCtrl {
	return &UserCtrl{
		service: service,
	}
}

func (ctrl *UserCtrl) SignUp(c *gin.Context) {
	//1.获取参数和参数校验
	var data = new(models.UserSignupParams)
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("c.ShouldBindJSON(data) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.业务处理
	if err := ctrl.service.Signup(data); err != nil {
		zap.L().Error("ctrl.service.Signup(data) failed", zap.Error(err))
		ResponseError(c, CodeUserExist)
		return
	}
	//3.返回响应
	ResponseSuccess(c, signupSuccess)
}

func (ctrl *UserCtrl) Login(c *gin.Context) {
	//1.获取参数并检验
	var data = new(models.UserLoginParams)
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("c.ShouldBindJSON(data) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.业务处理
	ret, err := ctrl.service.Login(data)
	if err != nil {
		zap.L().Error("ctrl.service.Login(data) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, ret)
}

func (ctrl *UserCtrl) Logout(c *gin.Context) {
	////1.参数验证 --> 得到相应的token
	//authHeader := c.Request.Header.Get("Authorization")
	//parts := strings.SplitN(authHeader, " ", 2)
	////得到token
	//token := parts[1]
	////2.业务处理 --> 将该token储存在数据库中
	//if err := ctrl.service.Logout(token); err != nil {
	//	zap.L().Error("logic.Logout(token) failed", zap.Error(err))
	//	ResponseError(c, CodeServeBusy)
	//	return
	//}
	////3.返回响应
	//ResponseSuccess(c, CodeSuccess)
}

func (ctrl *UserCtrl) Update(c *gin.Context) {
	//1.参数校验
	var data = new(models.UserUpdateParams)
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("c.ShouldBindJSON(data) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//获取id
	uid, err := getUserId(c)
	if err != nil {
		zap.L().Error("getUserId(c) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	data.UID = uid

	//2.业务处理
	if err = ctrl.service.Update(data); err != nil {
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
