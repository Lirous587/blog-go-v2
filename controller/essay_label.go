package controller

import (
	"blog/models"
	"blog/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type EssayLabelCtrl struct {
	service service.EssayLabelService
}

func NewEssayLabelCtrl(service service.EssayLabelService) *EssayLabelCtrl {
	return &EssayLabelCtrl{
		service: service,
	}
}

func (ctrl *EssayLabelCtrl) Create(c *gin.Context) {
	//1.参数处理
	var data = new(models.EssayLabelParam)
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("c.ShouldBindJSON(data) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.业务处理
	if err := ctrl.service.Create(data); err != nil {
		zap.L().Error("ctrl.service.Create(data) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, createSuccess)
}

func (ctrl *EssayLabelCtrl) Delete(c *gin.Context) {
	//1.参数处理
	qid := c.Query("id")
	id64, err := strconv.ParseInt(qid, 10, 64)
	if err != nil {
		zap.L().Error("parse id failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	id := int(id64)
	//2.业务处理
	if err := ctrl.service.Delete(id); err != nil {
		zap.L().Error("ctrl.service.Delete(id) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, deleteSuccess)
}

func (ctrl *EssayLabelCtrl) Update(c *gin.Context) {
	//1.参数处理
	var data = new(models.EssayLabelUpdateParam)
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("c.ShouldBindJSON(data) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.业务处理
	if err := ctrl.service.Update(data); err != nil {
		zap.L().Error("ctrl.service.Update(data) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, updateSuccess)
}
