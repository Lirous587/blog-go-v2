package logic

import (
	"blog/dao/mysql"
	"blog/models"
)

func CreateLabel(p *models.LabelParams) (err error) {
	return mysql.CreateLabel(p)
}

func DeleteLabel(id int) (err error) {
	return mysql.DeleteLabel(id)
}

func UpdateLabel(p *models.LabelUpdateParams) (err error) {
	return mysql.UpdateLabel(p)
}
