package service

import (
	"blog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type ImgService interface {
	Upload(*gin.Context, *multipart.FileHeader) (string, error)
}

type ImgLocalService struct {
}

func NewImgLocalService() *ImgLocalService {
	return &ImgLocalService{}
}

func (s *ImgLocalService) Upload(c *gin.Context, f *multipart.FileHeader) (string, error) {
	sanitizedFileName := utils.SanitizedFileName(f.Filename)
	// 将获取的文件保存到本地
	dst := fmt.Sprintf("/app/statics/img/%s", sanitizedFileName)
	if err := c.SaveUploadedFile(f, dst); err != nil {
		return "", err
	}
	return sanitizedFileName, nil
}
