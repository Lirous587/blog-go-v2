package controller

import (
	"blog/logic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

const (
	createHeartWordsSuccess = "创建心语成功"
	deleteHeartWordsSuccess = "删除心语成功"
	updateHeartWordsSuccess = "修改心语成功"
)

func CreateHeartWordsHandler(c *gin.Context) {
	p := new(models.HeartWordsParams)
	// 1.参数绑定
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	// 2.逻辑处理
	if err := logic.CreateHeartWords(p); err != nil {
		zap.L().Error("logic.CreateHeartWords(p) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, createHeartWordsSuccess)
}

func DeleteHeartWordsHandler(c *gin.Context) {
	//1.获取参数
	qid := c.Query("id")
	id, err := strconv.Atoi(qid)
	if err != nil {
		zap.L().Error("strconv.Atoi(qid) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.逻辑处理
	if err = logic.DeleteHeartWords(id); err != nil {
		zap.L().Error("logic.DeleteHeartWords(qid) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, deleteHeartWordsSuccess)
}

func UpdateHeartWordsHandler(c *gin.Context) {
	//1.参数检验
	var p = new(models.HeartWordsUpdateParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.逻辑处理
	if err := logic.UpdateHeartWords(p); err != nil {
		zap.L().Error("logic.UpdateHeartWords(p) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, updateHeartWordsSuccess)
}

func ResponseHeardWordsListHandler(c *gin.Context) {
	query := models.HeartWordsQuery{}
	page64, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	if page64 == 0 {
		page64 = 1
	}
	query.Page = int(page64)

	pageSize64, _ := strconv.ParseInt(c.Query("pageSize"), 10, 64)
	if pageSize64 == 0 {
		pageSize64 = 10
	}
	query.PageSize = int(pageSize64)

	ifCouldType := c.Query("couldType")
	if ifCouldType != "" {
		query.IfCouldType = true
	}

	var listAndPage = new(models.HeartWordsListAndPage)
	if err := logic.GetHeartWordsList(listAndPage, query); err != nil {
		zap.L().Error(" logic.GetHeartWordsList(listAndPage, query) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//得到各大分类种类以及相应的名称
	ResponseSuccess(c, listAndPage)
}
