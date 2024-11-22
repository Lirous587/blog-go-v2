package controller

import (
	"blog/models"
	"blog/server"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func newHeartWordsData() *models.HeartWordsData {
	return &models.HeartWordsData{}
}

func CreateHeartWordsHandler(c *gin.Context) {
	data := newHeartWordsData()
	obj := server.NewHeartWordsServer(data)
	// 1.参数绑定
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("c.ShouldBindJSON(model) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	// 2.逻辑处理
	if err := obj.Create(data); err != nil {
		zap.L().Error("service.Create() failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, createSuccess)
}

func DeleteHeartWordsHandler(c *gin.Context) {
	data := newHeartWordsData()
	obj := server.NewHeartWordsServer(data)
	//1.获取参数
	qid := c.Query("id")
	id, err := strconv.Atoi(qid)
	if err != nil {
		zap.L().Error("strconv.Atoi(qid) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	data.ID = id
	//2.逻辑处理
	if err = obj.Delete(data.ID); err != nil {
		zap.L().Error("service.Delete() failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, createSuccess)
}

func UpdateHeartWordsHandler(c *gin.Context) {
	data := newHeartWordsData()
	obj := server.NewHeartWordsServer(data)
	//1.参数检验
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("c.ShouldBindJSON(data) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.逻辑处理
	if err := obj.Update(data); err != nil {
		zap.L().Error("service.Update() failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, createSuccess)
}

func ResponseHeardWordsListHandler(c *gin.Context) {
	data := new(models.HeartWordsListAndPage)
	obj := server.NewHeartWordsServer(data)
	//	参数处理
	page := utils.DisposePageQuery(c)
	query := models.HeartWordsQuery{
		Page:     page.Page,
		PageSize: page.PageSize,
	}
	ifCouldType := c.Query("couldType")
	if ifCouldType != "" {
		query.IfCouldType = true
	}

	list, err := obj.GetList(query)
	if err != nil {
		zap.L().Error("obj.GetList(query) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//得到各大分类种类以及相应的名称
	ResponseSuccess(c, list)
}
