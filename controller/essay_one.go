package controller

import (
	"blog/logic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

const (
	createEssaySuccess = "添加文章成功"
	deleteEssaySuccess = "删除文章成功"
	updateEssaySuccess = "修改文章成功"
)

func ResponseEssayDataHandler(c *gin.Context) {
	//1.参数处理
	qid := c.Query("id")
	id, err := strconv.Atoi(qid)
	if err != nil {
		zap.L().Error("strconv.Atoi(qid) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	//2.业务处理
	var essay = new(models.EssayContent)
	essay.Id = id
	if err = logic.GetEssayData(essay); err != nil {
		zap.L().Error("logic.GetEssayData(essay, id) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, essay)
}

func CreateEssayHandler(c *gin.Context) {
	//1.参数处理
	var p = new(models.EssayParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//2.业务处理
	if err := logic.CreateEssay(p); err != nil {
		zap.L().Error("logic.CreateEssay(essay) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, createEssaySuccess)
}

func DeleteEssayHandler(c *gin.Context) {
	//1.获取参数
	qid := c.Query("id")
	id, err := strconv.Atoi(qid)
	if err != nil {
		zap.L().Error("strconv.Atoi(qid) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.逻辑处理
	if err = logic.DeleteEssay(id); err != nil {
		zap.L().Error("logic.DeleteEssay(id) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, deleteEssaySuccess)
}

func UpdateEssayHandler(c *gin.Context) {
	//1.获取参数
	var p = new(models.EssayUpdateParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}

	//2.业务处理
	if err := logic.UpdateEssay(p); err != nil {
		zap.L().Error("logic.UpdateEssay(p) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, updateEssaySuccess)
}
