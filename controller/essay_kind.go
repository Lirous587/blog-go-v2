package controller

import (
	"blog/models"
	"blog/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type EssayKindController struct {
	service service.EssayKindService
}

func NewEssayKindController(service service.EssayKindService) *EssayKindController {
	return &EssayKindController{
		service: service,
	}
}

func (ctrl *EssayKindController) Create(c *gin.Context) {
	data := new(models.EssayKindData)
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

func (ctrl *EssayKindController) Update(c *gin.Context) {
	//1.参数检验
	var data = new(models.EssayKindData)
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

func (ctrl *EssayKindController) Delete(c *gin.Context) {
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
