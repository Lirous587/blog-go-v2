package controller

import (
	"blog/models"
	"blog/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type GalleryKindController struct {
	service service.GalleryKindService
}

func NewGalleryKindController(service service.GalleryKindService) *GalleryKindController {
	return &GalleryKindController{
		service: service,
	}
}

func (ctrl *GalleryKindController) Create(c *gin.Context) {
	data := new(models.GalleryKindData)
	// 1.参数绑定
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("c.ShouldBindJSON(data) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}

	// 2.逻辑处理
	if err := ctrl.service.Create(data); err != nil {
		zap.L().Error("ctrl.service.Create(data) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	ResponseSuccess(c, createSuccess)
}

func (ctrl *GalleryKindController) Delete(c *gin.Context) {
	// 1.获取参数
	qid := c.Query("id")
	id, err := strconv.Atoi(qid)
	if err != nil {
		zap.L().Error("strconv.Atoi(qid) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}

	// 2.逻辑处理
	if err := ctrl.service.Delete(id); err != nil {
		zap.L().Error("ctrl.service.Delete(id) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	// 3.返回响应
	ResponseSuccess(c, deleteSuccess)
}

func (ctrl *GalleryKindController) Update(c *gin.Context) {
	data := new(models.GalleryKindData)
	// 1.参数检验
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("c.ShouldBindJSON(data) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}

	// 2.逻辑处理
	if err := ctrl.service.Update(data); err != nil {
		zap.L().Error("ctrl.service.Update(data) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	// 3.返回响应
	ResponseSuccess(c, updateSuccess)
}

func (ctrl *GalleryKindController) GetList(c *gin.Context) {
	list, err := ctrl.service.GetList()
	if err != nil {
		zap.L().Error("ctrl.service.GetList() failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//得到各大分类种类以及相应的名称
	ResponseSuccess(c, list)
}
