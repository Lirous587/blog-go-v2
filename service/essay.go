package service

import (
	"blog/cache"
	"blog/models"
	"blog/repository"
	"blog/utils"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"strings"
)

const (
	invalidLabelIds = "无效的标签"
)

//var essayAndKeywords []models.EssayDesc

type EssayService interface {
	Create(data *models.EssayParams) error
	Read(id int) (*models.EssayData, error)
	Update(data *models.EssayUpdateParams) error
	Delete(id int) error
	GetList(q *models.EssayQuery) (*models.EssayListAndPage, error)
	GetListBySearch(p *models.SearchParam) ([]models.EssayDesc, error)
	UpdateDescCache() error
}

type EssayRepoService struct {
	repo repository.EssayRepo
	cch  cache.EssayCache
}

func NewEssayRepoService(cch cache.EssayCache, repo repository.EssayRepo) *EssayRepoService {
	return &EssayRepoService{
		cch:  cch,
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
	return
}

func (s *EssayRepoService) Delete(id int) (err error) {
	return s.repo.Delete(id)
}

func (s *EssayRepoService) GetList(q *models.EssayQuery) (data *models.EssayListAndPage, err error) {
	data, err = s.repo.GetList(q)
	return
}

func (s *EssayRepoService) GetListBySearch(p *models.SearchParam) ([]models.EssayDesc, error) {
	allDesc, err := s.getAllDesc()
	if err != nil {
		return nil, fmt.Errorf("allDesc,err:= s.getAllDesc() failed,err:%w", err)
	}
	list := make([]models.EssayDesc, 0, 5)
	for _, essay := range allDesc {
		// 检查 essay.keyword 数组中是否包含指定的关键字 k
		if strings.Contains(essay.Keywords, p.Keyword) {
			list = append(list, essay)
		}
	}
	return list, nil
}

func (s *EssayRepoService) getAllDesc() ([]models.EssayDesc, error) {
	// 先从redis去查
	list, err := s.cch.GetDesc()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("s.cch.GetDesc() failed,err%w", err)
		}
		// 没有就去mysql去查 并且保证到redis中
		list, err := s.repo.GetAllDesc()
		if err != nil {
			return nil, fmt.Errorf("s.repo.GetAllDesc() failed,err%w", err)
		}
		err = s.cch.SaveDesc(list)
		if err != nil {
			return nil, fmt.Errorf("s.cch.SaveDesc(list) failed,err%w", err)
		}
		return list, nil
	}
	return list, nil
}

func (s *EssayRepoService) UpdateDescCache() error {
	// 删除cch
	if err := s.cch.DeleteDesc(); err != nil {
		return err
	}
	// 从repo读取
	data, err := s.repo.GetAllDesc()
	if err != nil {
		return err
	}
	// 保存到cch
	return s.cch.SaveDesc(data)
}
