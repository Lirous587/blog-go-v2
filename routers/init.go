package routers

import (
	"blog/controller"
	"blog/dao/mysql"
	"blog/repository"
	"blog/server"
)

func InitHeartWordsController() *controller.HeartWordsController {
	// 初始化仓库和服务
	var repo repository.HeartWordsRepo
	repo = repository.NewHeartWordsRepoMySQL(mysql.DB)
	var ser server.HeartWordsServer
	ser = server.NewHeartWordsRepoService(repo)
	// 初始化控制器
	return controller.NewHeartWordsController(ser)
}

func InitGalleryController() *controller.GalleryController {
	// 初始化仓库和服务
	var repo repository.GalleryRepo
	repo = repository.NewGalleryRepoMySQL(mysql.DB)
	var ser server.GalleryServer
	ser = server.NewGalleryRepoService(repo)
	// 初始化控制器
	return controller.NewGalleryController(ser)
}

func InitGalleryKindController() *controller.GalleryKindController {
	// 初始化仓库和服务
	var repo repository.GalleryKindRepo
	repo = repository.NewGalleryKindRepoMySQL(mysql.DB)
	var ser server.GalleryKindServer
	ser = server.NewGalleryKindRepoService(repo)
	// 初始化控制器
	return controller.NewGalleryKindController(ser)
}

func InitEssayKindController() *controller.EssayKindController {
	// 初始化仓库和服务
	var repo repository.EssayKindRepo
	repo = repository.NewEssayKindRepoMySQL(mysql.DB)
	var ser server.EssayKindServer
	ser = server.NewEssayKindRepoService(repo)
	// 初始化控制器
	return controller.NewEssayKindController(ser)
}
