package routers

import (
	"blog/cache"
	"blog/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	Init()

	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DisableConsoleColor()

	r := gin.Default()
	// 创建新的CORS中间件
	config := cors.DefaultConfig()
	//这里要设置端口的 前端是:80不用显示调用
	//config.AllowOrigins = []string{"https://Lirous.com", "https://www.Lirous.com", "http://localhost:3000", "http://localhost:3001"}
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	r.Use(cors.New(config))
	r.Static("/api/img", "/app/statics/img")
	r.Static("/api/file", "/app/statics/file")

	// 初始化控制器
	heartWordsCtl := InitHeartWordsCtrl()
	galleryCtl := InitGalleryCtrl()
	galleryKindCtl := InitGalleryKindCtrl()
	essayKindCtrl := InitEssayKindCtrl()
	essayLabelCtrl := InitEssayLabelCtrl()
	essayCtrl := InitEssayCtrl()
	indexCtrl := InitIndexCtrl()
	imgCtrl := InitImgCtrl()
	userCtrl := InitUserCtrl()
	uCch := cache.UserCache(cache.NewUserCacheRedis(client))
	userAuthMiddleware := middlewares.NewUserAuthMiddleware(uCch)
	managerAuthMiddleware := middlewares.NewManagerAuthMiddleware()

	indexGroup := r.Group("/api/base")
	{
		indexGroup.GET("/index", middlewares.SaveUserIp(), indexCtrl.GetData)
		indexGroup.GET("/essayList", essayCtrl.GetList)
		indexGroup.GET("/essayContent", essayCtrl.Read)
		indexGroup.POST("/essaySearch", essayCtrl.GetListBySearch)
		indexGroup.GET("/heartWordsList", heartWordsCtl.GetList)
	}

	userGroup := r.Group("/api/user")
	{
		userGroup.POST("/login", userCtrl.Login)
		userGroup.POST("/signup", userCtrl.SignUp)
		userGroup.POST("/logout", userAuthMiddleware, userCtrl.Logout)
		userGroup.PUT("/update", userAuthMiddleware, userCtrl.Update)
	}

	managerGroup := r.Group("/api/admin")
	managerGroup.Use(managerAuthMiddleware)

	{
		//managerGroup.POST("/login", managerCtrl.Login)
		// 上传图片
		managerGroup.POST("/uploadImg", imgCtrl.Upload)

		// 主页数据
		//managerGroup.GET("/panel", controller.ResponseDataAboutManagerPanel)

		//gallery
		managerGroup.GET("/galleryList", galleryCtl.GetList)
		managerGroup.POST("/gallery", galleryCtl.Create)
		managerGroup.DELETE("/gallery", galleryCtl.Delete)
		managerGroup.PUT("/gallery", galleryCtl.Update)

		//galleryKind
		managerGroup.GET("/galleryKindList", galleryKindCtl.GetList)
		managerGroup.POST("/galleryKind", galleryKindCtl.Create)
		managerGroup.DELETE("/galleryKind", galleryKindCtl.Delete)
		managerGroup.PUT("/galleryKind", galleryKindCtl.Update)
	}

	managerGroupIndex := r.Group("/api/admin")
	managerGroupIndex.Use(managerAuthMiddleware, middlewares.UpdateIndexMiddleware(indexCtrl))
	{
		// kind
		managerGroupIndex.POST("/kind", essayKindCtrl.Create)
		managerGroupIndex.DELETE("/kind", essayKindCtrl.Delete)
		managerGroupIndex.PUT("/kind", essayKindCtrl.Update)

		// label
		managerGroupIndex.POST("/label", essayLabelCtrl.Create)
		managerGroupIndex.DELETE("label", essayLabelCtrl.Delete)
		managerGroupIndex.PUT("/label", essayLabelCtrl.Update)

		//heartWord
		managerGroupIndex.POST("/heartWords", heartWordsCtl.Create)
		managerGroupIndex.DELETE("/heartWords", heartWordsCtl.Delete)
		managerGroupIndex.PUT("/heartWords", heartWordsCtl.Update)
	}

	managerGroupEssay := r.Group("/api/admin")
	managerGroupEssay.Use(managerAuthMiddleware, middlewares.UpdateIndexMiddleware(indexCtrl), middlewares.UpdateEssayDescMiddleware(essayCtrl))
	{
		// essay
		managerGroupEssay.POST("/essay", essayCtrl.Create)
		managerGroupEssay.DELETE("/essay", essayCtrl.Delete)
		managerGroupEssay.PUT("/essay", essayCtrl.Update)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSONP(404, gin.H{
			"msg": "404",
		})
	})
	return r
}
