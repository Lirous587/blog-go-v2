package controller

import (
	"blog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UploadImgHandler(c *gin.Context) {
	f, err := c.FormFile("img") //name值与input 中的name 一致
	if err != nil {
		zap.L().Error("c.FormFile(\"img\") failed err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	sanitizedFileName := utils.SanitizedFileName(f.Filename)
	// 将获取的文件保存到本地
	dst := fmt.Sprintf("/app/statics/img/%s", sanitizedFileName)
	if err := c.SaveUploadedFile(f, dst); err != nil {
		zap.L().Error("c.SaveUploadedFile(f, dst) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	ResponseSuccess(c, sanitizedFileName)
}
