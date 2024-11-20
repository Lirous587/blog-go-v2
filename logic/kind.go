package logic

import (
	"blog/dao/mysql"
	"blog/models"
)

func CreateKind(p *models.KindParams) error {
	return mysql.CreateKind(p)
}

func DeleteKind(id int) error {
	return mysql.DeleteKind(id)
}

func UpdateKind(p *models.KindUpdateParams) error {
	return mysql.UpdateKind(p)
}
