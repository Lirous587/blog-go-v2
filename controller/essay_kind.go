package controller

import (
	"blog/models"
	"blog/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type EssayKindCtrl struct {
	service service.EssayKindService
}

func NewEssayKindCtrl(service service.EssayKindService) *EssayKindCtrl {
	return &EssayKindCtrl{
		service: service,
	}
}

func (ctrl *EssayKindCtrl) Create(c *gin.Context) {
	data := new(models.EssayKindParam)
	// 1.参数绑定
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("c.ShouldBindJSON(data) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	// 2.逻辑处理
	if err := ctrl.service.Create(data); err != nil {
		zap.L().Error("ctrl.service.Create(data) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, createSuccess)
}

func (ctrl *EssayKindCtrl) Update(c *gin.Context) {
	//1.参数检验
	var data = new(models.EssayKindUpdateParam)
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("c.ShouldBindJSON(data) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.逻辑处理
	if err := ctrl.service.Update(data); err != nil {
		zap.L().Error("ctrl.service.Update(data) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, updateSuccess)
}

func (ctrl *EssayKindCtrl) Delete(c *gin.Context) {
	//1.获取参数
	qid := c.Query("id")
	id, err := strconv.Atoi(qid)
	if err != nil {
		zap.L().Error("strconv.Atoi(qid) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.逻辑处理
	if err = ctrl.service.Delete(id); err != nil {
		zap.L().Error("ctrl.service.Delete(id) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, deleteSuccess)
}
