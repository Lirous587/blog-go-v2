package controller

import (
	"blog/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IndexCtrl struct {
	service service.IndexService
}

func NewIndexCtrl(service service.IndexService) *IndexCtrl {
	return &IndexCtrl{
		service: service,
	}
}

func (ctrl *IndexCtrl) GetData(c *gin.Context) {
	data, err := ctrl.service.GetData()
	if err != nil {
		zap.L().Error("ctrl.service.GetData() failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, data)
}

func (ctrl *IndexCtrl) Update() error {
	err := ctrl.service.Update()
	if err != nil {
		zap.L().Error("ctrl.service.Update() failed", zap.Error(err))
	}
	return err
}
