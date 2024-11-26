package routers

import (
	"blog/controller"
	"blog/dao/mysql"
	"blog/repository"
	"blog/service"
)

func InitHeartWordsController() *controller.HeartWordsController {
	// 初始化仓库和服务
	var repo repository.HeartWordsRepo
	repo = repository.NewHeartWordsRepoMySQL(mysql.DB)
	var ser service.HeartWordsService
	ser = service.NewHeartWordsRepoService(repo)
	// 初始化控制器
	return controller.NewHeartWordsController(ser)
}

func InitGalleryController() *controller.GalleryController {
	// 初始化仓库和服务
	var repo repository.GalleryRepo
	repo = repository.NewGalleryRepoMySQL(mysql.DB)
	var ser service.GalleryService
	ser = service.NewGalleryRepoService(repo)
	// 初始化控制器
	return controller.NewGalleryController(ser)
}

func InitGalleryKindController() *controller.GalleryKindController {
	// 初始化仓库和服务
	var repo repository.GalleryKindRepo
	repo = repository.NewGalleryKindRepoMySQL(mysql.DB)
	var ser service.GalleryKindService
	ser = service.NewGalleryKindRepoService(repo)
	// 初始化控制器
	return controller.NewGalleryKindController(ser)
}

func InitEssayKindController() *controller.EssayKindController {
	// 初始化仓库和服务
	var repo repository.EssayKindRepo
	repo = repository.NewEssayKindRepoMySQL(mysql.DB)
	var ser service.EssayKindService
	ser = service.NewEssayKindRepoService(repo)
	// 初始化控制器
	return controller.NewEssayKindController(ser)
}

func InitEssayLabelController() *controller.EssayLabelCtrl {
	// 初始化仓库和服务
	var repo repository.EssayLabelRepo
	repo = repository.NewEssayLabelRepoMySQL(mysql.DB)
	var ser service.EssayLabelService
	ser = service.NewEssayLabelRepoService(repo)
	// 初始化控制器
	return controller.NewEssayLabelController(ser)
}
