package repository

import (
	"blog/models"
	"github.com/jmoiron/sqlx"
)

type IndexRepo interface {
	GetData() (*models.IndexData, error)
}

type IndexRepoMySQL struct {
	db *sqlx.DB
}

func NewIndexRepoMySql(db *sqlx.DB) *IndexRepoMySQL {
	return &IndexRepoMySQL{
		db: db,
	}
}

func (r *IndexRepoMySQL) GetData() (data *models.IndexData, err error) {
	ekRepo := EssayKindRepo(NewEssayKindRepoMySQL(r.db))
	kindList, err := ekRepo.GetList()
	if err != nil {
		return
	}

	elRepo := EssayLabelRepo(NewEssayLabelRepoMySQL(r.db))
	labelList, err := elRepo.GetList()
	if err != nil {
		return
	}

	eRepo := EssayRepo(NewEssayRepoMySQL(r.db))
	essayList, err := eRepo.GetRecommendList()
	if err != nil {
		return
	}

	hwRepo := HeartWordsRepo(NewHeartWordsRepoMySQL(r.db))
	heartWordsList, err := hwRepo.GetCouldTypeList()
	if err != nil {
		return
	}
	data = new(models.IndexData)
	//整合数据
	data.KindList = kindList
	data.LabelList = labelList
	data.EssayList = essayList
	data.HeartWordsList = heartWordsList
	return
}
