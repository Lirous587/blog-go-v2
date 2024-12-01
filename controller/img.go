package controller

import (
	"blog/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ImgCtrl struct {
	service service.ImgService
}

func NewImgCtrl(service service.ImgService) *ImgCtrl {
	return &ImgCtrl{
		service: service,
	}
}

func (ctrl *ImgCtrl) Upload(c *gin.Context) {
	f, err := c.FormFile("img") //name值与input 中的name 一致
	if err != nil {
		zap.L().Error("c.FormFile(\"img\") failed err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	sanitizedFileName, err := ctrl.service.Upload(c, f)
	if err != nil {
		zap.L().Error("ctrl.service.Upload() failed err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, sanitizedFileName)
}
