package routers

import (
	"blog/cache"
	"blog/controller"
	"blog/repository"
	"blog/service"
)

func InitHeartWordsCtrl() *controller.HeartWordsCtrl {
	// 初始化仓库和服务
	var repo repository.HeartWordsRepo
	repo = repository.NewHeartWordsRepoMySQL(repository.DB)
	var ser service.HeartWordsService
	ser = service.NewHeartWordsRepoService(repo)
	// 初始化控制器
	return controller.NewHeartWordsCtrl(ser)
}

func InitGalleryCtrl() *controller.GalleryCtrl {
	// 初始化仓库和服务
	var repo repository.GalleryRepo
	repo = repository.NewGalleryRepoMySQL(repository.DB)
	var ser service.GalleryService
	ser = service.NewGalleryRepoService(repo)
	// 初始化控制器
	return controller.NewGalleryCtrl(ser)
}

func InitGalleryKindCtrl() *controller.GalleryKindCtrl {
	// 初始化仓库和服务
	var repo repository.GalleryKindRepo
	repo = repository.NewGalleryKindRepoMySQL(repository.DB)
	var ser service.GalleryKindService
	ser = service.NewGalleryKindRepoService(repo)
	// 初始化控制器
	return controller.NewGalleryKindCtrl(ser)
}

func InitEssayKindCtrl() *controller.EssayKindCtrl {
	// 初始化仓库和服务
	var repo repository.EssayKindRepo
	repo = repository.NewEssayKindRepoMySQL(repository.DB)
	var ser service.EssayKindService
	ser = service.NewEssayKindRepoService(repo)
	// 初始化控制器
	return controller.NewEssayKindCtrl(ser)
}

func InitEssayLabelCtrl() *controller.EssayLabelCtrl {
	// 初始化仓库和服务
	var repo repository.EssayLabelRepo
	repo = repository.NewEssayLabelRepoMySQL(repository.DB)
	var ser service.EssayLabelService
	ser = service.NewEssayLabelRepoService(repo)
	// 初始化控制器
	return controller.NewEssayLabelCtrl(ser)
}

func InitEssayCtrl() *controller.EssayCtrl {
	// 初始化仓库和服务
	var repo repository.EssayRepo
	repo = repository.NewEssayRepoMySQL(repository.DB)
	var ser service.EssayService
	ser = service.NewEssayRepoService(repo)
	// 初始化控制器
	return controller.NewEssayCtrl(ser)
}

func InitIndexCtrl() *controller.IndexCtrl {
	// 初始化仓库和服务
	var cch cache.IndexCache
	cch = cache.NewIndexCacheRedis(cache.Rdb)
	var ser service.IndexService
	var repo repository.IndexRepo
	repo = repository.NewIndexRepoMySql(repository.DB)
	ser = service.NewIndexDataCacheService(cch, repo)
	// 初始化控制器
	return controller.NewIndexCtrl(ser)
}
