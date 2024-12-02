package routers

import (
	"blog/cache"
	"blog/controller"
	"blog/repository"
	"blog/service"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

var (
	db     *sqlx.DB
	client *redis.Client
)

func Init() {
	var err error
	db, err = repository.MysqlInit()
	if err != nil {
		panic(err)
	}
	client, err = cache.RedisInit()
	if err != nil {
		panic(err)
	}
}

func InitHeartWordsCtrl() *controller.HeartWordsCtrl {
	// 初始化仓库和服务
	var repo repository.HeartWordsRepo
	repo = repository.NewHeartWordsRepoMySQL(db)
	var ser service.HeartWordsService
	ser = service.NewHeartWordsRepoService(repo)
	// 初始化控制器
	return controller.NewHeartWordsCtrl(ser)
}

func InitGalleryCtrl() *controller.GalleryCtrl {
	// 初始化仓库和服务
	var repo repository.GalleryRepo
	repo = repository.NewGalleryRepoMySQL(db)
	var ser service.GalleryService
	ser = service.NewGalleryRepoService(repo)
	// 初始化控制器
	return controller.NewGalleryCtrl(ser)
}

func InitGalleryKindCtrl() *controller.GalleryKindCtrl {
	// 初始化仓库和服务
	var repo repository.GalleryKindRepo
	repo = repository.NewGalleryKindRepoMySQL(db)
	var ser service.GalleryKindService
	ser = service.NewGalleryKindRepoService(repo)
	// 初始化控制器
	return controller.NewGalleryKindCtrl(ser)
}

func InitEssayKindCtrl() *controller.EssayKindCtrl {
	// 初始化仓库和服务
	var repo repository.EssayKindRepo
	repo = repository.NewEssayKindRepoMySQL(db)
	var ser service.EssayKindService
	ser = service.NewEssayKindRepoService(repo)
	// 初始化控制器
	return controller.NewEssayKindCtrl(ser)
}

func InitEssayLabelCtrl() *controller.EssayLabelCtrl {
	// 初始化仓库和服务
	var repo repository.EssayLabelRepo
	repo = repository.NewEssayLabelRepoMySQL(db)
	var ser service.EssayLabelService
	ser = service.NewEssayLabelRepoService(repo)
	// 初始化控制器
	return controller.NewEssayLabelCtrl(ser)
}

func InitEssayCtrl() *controller.EssayCtrl {
	// 初始化仓库和服务
	var cch cache.EssayCache
	cch = cache.NewEssayCacheRedis(client)
	var repo repository.EssayRepo
	repo = repository.NewEssayRepoMySQL(db)
	var ser service.EssayService
	ser = service.NewEssayRepoService(cch, repo)
	// 初始化控制器
	return controller.NewEssayCtrl(ser)
}

func InitIndexCtrl() *controller.IndexCtrl {
	// 初始化仓库和服务
	var cch cache.IndexCache
	cch = cache.NewIndexCacheRedis(client)
	var ser service.IndexService
	var repo repository.IndexRepo
	repo = repository.NewIndexRepoMySql(db)
	ser = service.NewIndexDataCacheService(cch, repo)
	// 初始化控制器
	return controller.NewIndexCtrl(ser)
}

func InitImgCtrl() *controller.ImgCtrl {
	var ser service.ImgService
	ser = service.NewImgLocalService()
	// 初始化控制器
	return controller.NewImgCtrl(ser)
}

func InitUserCtrl() *controller.ImgCtrl {
	var ser service.ImgService
	ser = service.NewImgLocalService()
	// 初始化控制器
	return controller.NewImgCtrl(ser)
}
