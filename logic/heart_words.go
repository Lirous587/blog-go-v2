package logic

import (
	"blog/dao/mysql"
	"blog/models"
)

func CreateHeartWords(p *models.HeartWordsParams) error {
	return mysql.CreateHeartWords(p)
}

func DeleteHeartWords(id int) error {
	return mysql.DeleteHeartWords(id)
}

func UpdateHeartWords(p *models.HeartWordsUpdateParams) error {
	return mysql.UpdateHeartWords(p)
}

func GetHeartWordsList(list *models.HeartWordsListAndPage, query models.HeartWordsQuery) error {
	if err := mysql.GetHeartWordsList(list, query); err != nil {
		return err
	}
	return nil
}
