package controller

import (
	"blog/models"
	"blog/server"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type HeartWordsController struct {
	server server.HeartWordsServer
}

func NewHeartWordsController(server server.HeartWordsServer) *HeartWordsController {
	return &HeartWordsController{
		server: server,
	}
}

func (ctrl *HeartWordsController) Create(c *gin.Context) {
	data := new(models.HeartWordsData)
	// 1.参数绑定
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("c.ShouldBindJSON(data) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}

	// 2.逻辑处理
	if err := ctrl.server.Create(data); err != nil {
		zap.L().Error("ctrl.service.Create(data) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	ResponseSuccess(c, createSuccess)
}

func (ctrl *HeartWordsController) Delete(c *gin.Context) {
	// 1.获取参数
	qid := c.Query("id")
	id, err := strconv.Atoi(qid)
	if err != nil {
		zap.L().Error("strconv.Atoi(qid) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}

	// 2.逻辑处理
	if err := ctrl.server.Delete(id); err != nil {
		zap.L().Error("ctrl.service.Delete(id) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	// 3.返回响应
	ResponseSuccess(c, deleteSuccess)
}

func (ctrl *HeartWordsController) Update(c *gin.Context) {
	data := new(models.HeartWordsData)
	// 1.参数检验
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("c.ShouldBindJSON(data) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}

	// 2.逻辑处理
	if err := ctrl.server.Update(data); err != nil {
		zap.L().Error("ctrl.service.Update(data) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	// 3.返回响应
	ResponseSuccess(c, updateSuccess)
}

func (ctrl *HeartWordsController) GetList(c *gin.Context) {
	//	参数处理
	query := new(models.HeartWordsQuery)
	if err := c.ShouldBindQuery(query); err != nil {
		zap.L().Error("c.ShouldBindQuery(query) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}

	list, err := ctrl.server.GetList(query)
	if err != nil {
		zap.L().Error("ctrl.service.GetList(query) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//得到各大分类种类以及相应的名称
	ResponseSuccess(c, list)
}
