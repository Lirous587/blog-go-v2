package controller

import (
	"blog/models"
	"blog/service"
	"github.com/gin-gonic/gin"
)

func ResponseIndexDataHandler(c *gin.Context) {
	var data = new(models.IndexData)
	service.GetIndexData(&data)
	ResponseSuccess(c, *data)
}
