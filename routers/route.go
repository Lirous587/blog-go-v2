package routers

import (
	"blog/controller"
	"blog/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
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

	v0 := r.Group("/api/base")
	v0.Use(middlewares.SaveUserIp())
	{
		v0.GET("/index", controller.ResponseIndexDataHandler)
	}

	v1 := r.Group("/api/base")
	{
		v1.GET("/essayList", essayCtrl.GetList)
		v1.GET("/essayContent", essayCtrl.Read)
		v1.GET("/heartWordsList", heartWordsCtl.GetList)
	}

	v2 := r.Group("/api/admin")
	{
		v2.POST("/login", controller.LoginHandler)
		//v2.POST("/signup", controller.SignupHandler)
		//v2.POST("/logout", controller.LogoutHandler)
		//v2.POST("/updateUserMsg", middlewares.JWTAuthMiddleware(), controller.UpdateUserMsgHandler)
	}

	v3Cache := r.Group("/api/admin")
	v3Cache.Use(middlewares.JWTAuthMiddleware(), middlewares.UpdateDataMiddleware())
	{
		// essay
		v3Cache.POST("/essay", essayCtrl.Create)
		v3Cache.DELETE("/essay", essayCtrl.Delete)
		v3Cache.PUT("/essay", essayCtrl.Update)

		// label
		v3Cache.POST("/label", essayLabelCtrl.Create)
		v3Cache.DELETE("label", essayLabelCtrl.Delete)
		v3Cache.PUT("/label", essayLabelCtrl.Update)

		// kind
		v3Cache.POST("/kind", essayKindCtrl.Create)
		v3Cache.DELETE("/kind", essayKindCtrl.Delete)
		v3Cache.PUT("/kind", essayKindCtrl.Update)
	}

	v3NoCache := r.Group("/api/admin")
	v3NoCache.Use(middlewares.JWTAuthMiddleware())
	{
		// 上传图片
		v3NoCache.POST("/uploadImg", controller.UploadImgHandler)

		// 主页数据
		//v3NoCache.GET("/panel", controller.ResponseDataAboutManagerPanel)

		//heartWord
		v3Cache.POST("/heartWords", heartWordsCtl.Create)
		v3Cache.DELETE("/heartWords", heartWordsCtl.Delete)
		v3Cache.PUT("/heartWords", heartWordsCtl.Update)

		//gallery
		v3NoCache.GET("/galleryList", galleryCtl.GetList)
		v3NoCache.POST("/gallery", galleryCtl.Create)
		v3NoCache.DELETE("/gallery", galleryCtl.Delete)
		v3NoCache.PUT("/gallery", galleryCtl.Update)

		//galleryKind
		v3NoCache.GET("/galleryKindList", galleryKindCtl.GetList)
		v3NoCache.POST("/galleryKind", galleryKindCtl.Create)
		v3NoCache.DELETE("/galleryKind", galleryKindCtl.Delete)
		v3NoCache.PUT("/galleryKind", galleryKindCtl.Update)
	}

	v4 := r.Group("/api/keyword")
	{
		v4.POST("/search", essayCtrl.GetListBySearch)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSONP(404, gin.H{
			"msg": "404",
		})
	})
	return r
}
