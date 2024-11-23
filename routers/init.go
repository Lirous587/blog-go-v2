package routers

import (
	"blog/controller"
	"blog/dao/mysql"
	"blog/server"
)

func InitHeartWordsController() *controller.HeartWordsController {
	// 初始化仓库和服务
	heartWordsRepo := mysql.NewHeartWordsMysql(mysql.DB)
	heartWordsService := server.NewHeartWordsServer(heartWordsRepo)
	// 初始化控制器
	return controller.NewHeartWordsController(heartWordsService)
}
