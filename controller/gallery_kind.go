package controller

import (
	"blog/logic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

const (
	createGalleryKindSuccess = "创建图片分类成功"
	deleteGalleryKindSuccess = "删除图片分类成功"
	updateGalleryKindSuccess = "修改图片分类成功"
)

func CreateGalleryKindHandler(c *gin.Context) {
	p := new(models.GalleryKindParams)
	// 1.参数绑定
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	// 2.逻辑处理
	if err := logic.CreateGalleryKind(p); err != nil {
		zap.L().Error("logic.CreateGalleryKind(p) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, createGalleryKindSuccess)
}

func DeleteGalleryKindHandler(c *gin.Context) {
	//1.获取参数
	qid := c.Query("id")
	id, err := strconv.Atoi(qid)
	if err != nil {
		zap.L().Error("strconv.Atoi(qid) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.逻辑处理
	if err = logic.DeleteGalleryKind(id); err != nil {
		zap.L().Error("logic.DeleteGalleryKind(id) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, deleteGalleryKindSuccess)
}

func UpdateGalleryKindHandler(c *gin.Context) {
	//1.参数检验
	var p = new(models.GalleryKindUpdateParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.逻辑处理
	if err := logic.UpdateGalleryKind(p); err != nil {
		zap.L().Error("logic.UpdateGalleryKind(p) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, updateGalleryKindSuccess)
}

func ResponseGalleryKindListHandler(c *gin.Context) {
	var list = new(models.GalleryKindList)
	if err := logic.GetGalleryKindList(list); err != nil {
		zap.L().Error("logic.GetGalleryList(listAndPage, query) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//得到各大分类种类以及相应的名称
	ResponseSuccess(c, list)
}
