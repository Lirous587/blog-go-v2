package controller

import (
	"blog/models"
	"blog/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type EssayCtrl struct {
	service service.EssayService
}

func NewEssayCtrl(service service.EssayService) *EssayCtrl {
	return &EssayCtrl{
		service: service,
	}
}

func (ctrl *EssayCtrl) Create(c *gin.Context) {
	//1.参数处理
	var data = new(models.EssayParams)
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("c.ShouldBindJSON(data) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//2.业务处理
	if err := ctrl.service.Create(data); err != nil {
		zap.L().Error("ctrl.service.Create(data) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, createSuccess)
}

func (ctrl *EssayCtrl) Read(c *gin.Context) {
	//1.参数处理
	qid := c.Query("id")
	id, err := strconv.Atoi(qid)
	if err != nil {
		zap.L().Error("strconv.Atoi(qid) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	//2.业务处理
	essay, err := ctrl.service.Read(id)
	if err != nil {
		zap.L().Error("ctrl.service.Read(id) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	essay.ID = id
	//3.返回响应
	ResponseSuccess(c, essay)
}

func (ctrl *EssayCtrl) Delete(c *gin.Context) {
	//1.获取参数
	qid := c.Query("id")
	id, err := strconv.Atoi(qid)
	if err != nil {
		zap.L().Error("strconv.Atoi(qid) failed,err:", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}
	//2.逻辑处理
	if err = ctrl.service.Delete(id); err != nil {
		zap.L().Error("ctrl.service.Delete(id) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, deleteSuccess)
}

func (ctrl *EssayCtrl) Update(c *gin.Context) {
	//1.获取参数
	var data = new(models.EssayUpdateParams)
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("c.ShouldBindJSON(data) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}

	//2.业务处理
	if err := ctrl.service.Update(data); err != nil {
		zap.L().Error("ctrl.service.Update(data) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, updateSuccess)
}

func (ctrl *EssayCtrl) GetList(c *gin.Context) {
	query := new(models.EssayQuery)
	if query.PageSize <= 0 || query.Page <= 0 {
		query.PageSize = 5
		query.Page = 1
	}
	if err := c.ShouldBind(query); err != nil {
		zap.L().Error("c.ShouldBind(query) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}

	listAndPage, err := ctrl.service.GetList(query)

	if err != nil {
		zap.L().Error("ctrl.service.GetList(query) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//得到各大分类种类以及相应的名称
	ResponseSuccess(c, listAndPage)
}

func (ctrl *EssayCtrl) GetListBySearch(c *gin.Context) {
	//1.参数检验
	p := new(models.SearchParam)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) failed", zap.Error(err))
		ResponseError(c, CodeParamInvalid)
		return
	}

	//2.逻辑处理
	essayList, err := ctrl.service.GetListBySearch(p)
	if err != nil {
		zap.L().Error("ctrl.service.GetListBySearch(p) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	//3.返回响应
	ResponseSuccess(c, essayList)
}

func (ctrl *EssayCtrl) UpdateDescCache() error {
	err := ctrl.service.UpdateDescCache()
	if err != nil {
		zap.L().Error("ctrl.service.Update() failed", zap.Error(err))
	}
	return err
}
