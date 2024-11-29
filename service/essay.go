package service

import (
	"blog/cache"
	"blog/models"
	"blog/repository"
	"blog/utils"
	"fmt"
	"strings"
)

const (
	invalidLabelIds = "无效的标签"
)

type EssayService interface {
	Create(data *models.EssayParams) error
	Read(id int) (*models.EssayData, error)
	Update(data *models.EssayUpdateParams) error
	Delete(id int) error
	GetList(q *models.EssayQuery) (*models.EssayListAndPage, error)
	GetListBySearch(p *models.SearchParam) ([]models.EssayData, error)
}

type EssayRepoService struct {
	repo repository.EssayRepo
}

func NewEssayRepoService(repo repository.EssayRepo) *EssayRepoService {
	return &EssayRepoService{
		repo: repo,
	}
}

func (s *EssayRepoService) Create(data *models.EssayParams) (err error) {
	if len(data.LabelIds) == 0 {
		return fmt.Errorf(invalidLabelIds)
	}
	if data.CreatedTime, err = utils.GetChineseTime(); err != nil {
		return fmt.Errorf("get chinese time failed: %w", err)
	}
	return s.repo.Create(data)
}

func (s *EssayRepoService) Read(id int) (data *models.EssayData, err error) {
	return s.repo.Read(id)
}

func (s *EssayRepoService) Update(data *models.EssayUpdateParams) (err error) {
	//更新数据
	if err = s.repo.Update(data); err != nil {
		return
	}
	idKeywords := new(models.EssayIdAndKeyword)
	idKeywords.EssayId = data.ID
	idKeywords.Keywords = data.Keywords
	if err = cache.SetEssayKeyword(idKeywords); err != nil {
		return
	}
	return
}

func (s *EssayRepoService) Delete(id int) (err error) {
	//删除redis中文章的相关数据
	if err = cache.DeleteEssay(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *EssayRepoService) GetList(q *models.EssayQuery) (data *models.EssayListAndPage, err error) {
	data, err = s.repo.GetList(q)
	return
}

func (s *EssayRepoService) GetListBySearch(p *models.SearchParam) (data []models.EssayData, err error) {
	//判断是否需要添加访问值
	if p.IfAdd {
		preKey := cache.KeySearchKeyWordTimes
		// 向redis中加入keyWord
		if err = cache.IncreaseSearchKeyword(preKey, (*p).Keyword); err != nil {
			return nil, err
		}
	}
	data = make([]models.EssayData, 0, 5)
	essayList := cache.GetAllEssayList()
	for _, essay := range essayList {
		// 检查 essay.keyword 数组中是否包含指定的关键字 k
		for _, keyword := range essay.Keywords {
			if strings.Contains(keyword, p.Keyword) {
				data = append(data, essay)
				break
			}
		}
	}
	return
}
