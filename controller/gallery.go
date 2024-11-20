package controller

import (
	"blog/logic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

const (
	createGallerySuccess = "创建图片成功"
	deleteGallerySuccess = "删除图片成功"
	updateGallerySuccess = "修改图片成功"
)

func CreateGalleryHandler(c *gin.Context) {
	p := new(models.GalleryParams)
	// 1.参数绑定
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	// 2.逻辑处理
	if err := logic.CreateGallery(p); err != nil {
		zap.L().Error("logic.CreateGallery(p) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, createGallerySuccess)
}

func DeleteGalleryHandler(c *gin.Context) {
	//1.获取参数
	qid := c.Query("id")
	id, err := strconv.Atoi(qid)
	if err != nil {
		zap.L().Error("strconv.Atoi(qid) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.逻辑处理
	if err = logic.DeleteGallery(id); err != nil {
		zap.L().Error("logic.DeleteGallery(id) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, deleteGallerySuccess)
}

func UpdateGalleryHandler(c *gin.Context) {
	//1.参数检验
	var p = new(models.GalleryUpdateParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.逻辑处理
	if err := logic.UpdateGallery(p); err != nil {
		zap.L().Error("logic.UpdateGallery(p) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, updateGallerySuccess)
}

func ResponseGalleryListHandler(c *gin.Context) {
	query := models.GalleryQuery{}
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

	kidSize64, _ := strconv.ParseInt(c.Query("kindID"), 10, 64)
	if kidSize64 == 0 {
		ResponseError(c, CodeParamInvalid)
		return
	}
	query.KindID = int(kidSize64)

	var listAndPage = new(models.GalleryListAndPage)
	if err := logic.GetGalleryList(listAndPage, query); err != nil {
		zap.L().Error("logic.GetGalleryList(listAndPage, query) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//得到各大分类种类以及相应的名称
	ResponseSuccess(c, listAndPage)
}
