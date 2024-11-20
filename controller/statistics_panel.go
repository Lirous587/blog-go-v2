package controller

import (
	"blog/logic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ResponseDataAboutManagerPanel(c *gin.Context) {
	list := new(models.Panel)

	if err := logic.GetUserIpCount(&list.IpSet); err != nil {
		zap.L().Error("logic.GetUserIpCount(&panelList.IpSet) failed,err", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	// 逻辑处理
	if err := logic.GetSearchKeywordRank(&list.RankZset); err != nil {
		zap.L().Error(" logic.GetSearchKeywordRank(&panelList.RankZset) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//	返回响应
	ResponseSuccess(c, list)
}
