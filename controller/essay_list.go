package controller

import (
	"blog/logic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func ResponseEssayListHandler(c *gin.Context) {
	query := models.EssayQuery{}
	page64, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	if page64 == 0 {
		page64 = 1
	}
	query.Page = int(page64)

	pageSize64, _ := strconv.ParseInt(c.Query("pageSize"), 10, 64)
	if pageSize64 == 0 {
		pageSize64 = 5
	}
	query.PageSize = int(pageSize64)

	lID, _ := strconv.ParseInt(c.Query("labelID"), 10, 64)
	KID, _ := strconv.ParseInt(c.Query("kindID"), 10, 64)

	query.LabelID = int(lID)
	query.KindID = int(KID)

	var listAndPage = new(models.EssayListAndPage)
	if err := logic.GetEssayList(listAndPage, query); err != nil {
		zap.L().Error("logic.GetDataAboutClassifyEssayMsg(listAndPage) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//得到各大分类种类以及相应的名称
	ResponseSuccess(c, listAndPage)
}

func ResponseDataAboutSearchKeyword(c *gin.Context) {
	//1.参数检验
	p := new(models.SearchParam)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}

	var essayList = new([]models.EssayData)
	//2.逻辑处理
	if err := logic.GetDataByKeyword(essayList, p); err != nil {
		zap.L().Error("logic.GetDataByKeyword(essayList, p) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	//3.返回响应
	ResponseSuccess(c, essayList)
}
