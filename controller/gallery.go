package controller

import (
	"blog/models"
	"blog/server"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type GalleryController struct {
	service server.GalleryServer
}

func NewGalleryController(service server.GalleryServer) *GalleryController {
	return &GalleryController{
		service: service,
	}
}

func (ctrl *GalleryController) Create(c *gin.Context) {
	data := new(models.GalleryData)
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

func (ctrl *GalleryController) Delete(c *gin.Context) {
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

func (ctrl *GalleryController) Update(c *gin.Context) {
	data := new(models.GalleryData)
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

func (ctrl *GalleryController) GetList(c *gin.Context) {
	//	参数处理
	page := utils.DisposePageQuery(c)
	query := models.GalleryQuery{
		Page:     page.Page,
		PageSize: page.PageSize,
	}

	kidSize64, _ := strconv.ParseInt(c.Query("kindID"), 10, 64)
	if kidSize64 == 0 {
		ResponseError(c, CodeParamInvalid)
		return
	}
	query.KindID = int(kidSize64)

	list, err := ctrl.service.GetList(query)
	if err != nil {
		zap.L().Error("ctrl.service.GetList(query) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//得到各大分类种类以及相应的名称
	ResponseSuccess(c, list)
}
