package server

import (
	"blog/dao/mysql"
	"blog/models"
)

type HeartWordsServer struct {
	data *models.HeartWordsData
}

type HeartWordsListServer struct {
	data *models.HeartWordsData
}

func NewHeartWordsServer(data *models.HeartWordsData) models.OneCruder[models.HeartWordsData] {
	return &HeartWordsServer{
		data: data,
	}
}

func NewHeartWordsListServer(data *models.HeartWordsData) *HeartWordsListServer {
	return &HeartWordsListServer{
		data: data,
	}
}

func (hws *HeartWordsServer) Create(data *models.HeartWordsData) error {
	obj := mysql.NewHeartWordsMysql(data)
	return obj.Create(data)
}

func (hws *HeartWordsServer) Read(id int) (data *models.HeartWordsData, err error) {
	obj := mysql.NewHeartWordsMysql(hws.data)
	return obj.Read(id)
}

func (hws *HeartWordsServer) Update(data *models.HeartWordsData) error {
	obj := mysql.NewHeartWordsMysql(data)
	return obj.Update(data)
}

func (hws *HeartWordsServer) Delete(id int) error {
	obj := mysql.NewHeartWordsMysql(hws.data)
	return obj.Delete(id)
}

func (hws *HeartWordsListServer) ReadList(query models.HeartWordsQuery) (*models.ListAndPage[models.HeartWordsData], error) {
	list := &models.ListAndPage[models.HeartWordsData]{}
	err := mysql.GetHeartWordsList(list, query)
	return list, err
}
