package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Page struct {
	Page     int
	PageSize int
}

func DisposePageQuery(c *gin.Context) Page {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	return Page{
		Page:     pageInt,
		PageSize: pageSizeInt,
	}

}
