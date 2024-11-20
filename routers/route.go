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

	v0 := r.Group("/api/base")
	v0.Use(middlewares.SaveUserIp())
	{
		v0.GET("/index", controller.ResponseIndexDataHandler)
	}

	v1 := r.Group("/api/base")
	{
		v1.GET("/essay_list", controller.ResponseEssayListHandler)
		v1.GET("/essay_content", controller.ResponseEssayDataHandler)
		v1.GET("/heartWords_list", controller.ResponseHeardWordsListHandler)
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
		v3Cache.POST("/essay", controller.CreateEssayHandler)
		v3Cache.DELETE("/essay", controller.DeleteEssayHandler)
		v3Cache.PUT("/essay", controller.UpdateEssayHandler)

		// label
		v3Cache.POST("/label", controller.CreateLabelHandler)
		v3Cache.DELETE("label", controller.DeleteLabelHandler)
		v3Cache.PUT("/label", controller.UpdateLabelHandler)

		// kind
		v3Cache.POST("/kind", controller.CreateKindHandler)
		v3Cache.DELETE("/kind", controller.DeleteKindHandler)
		v3Cache.PUT("/kind", controller.UpdateKindHandler)

		//heartWord
		v3Cache.POST("/heartWords", controller.CreateHeartWordsHandler)
		v3Cache.DELETE("/heartWords", controller.DeleteHeartWordsHandler)
		v3Cache.PUT("/heartWords", controller.UpdateHeartWordsHandler)
	}

	v3NoCache := r.Group("/api/admin")
	v3NoCache.Use(middlewares.JWTAuthMiddleware())
	{
		// 上传图片
		v3NoCache.POST("/uploadImg", controller.UploadImgHandler)
		// 主页数据
		v3NoCache.GET("/panel", controller.ResponseDataAboutManagerPanel)

		//gallery
		v3NoCache.GET("/gallery_list", controller.ResponseGalleryListHandler)
		v3NoCache.POST("/gallery", controller.CreateGalleryHandler)
		v3NoCache.DELETE("/gallery", controller.DeleteGalleryHandler)
		v3NoCache.PUT("/gallery", controller.UpdateGalleryHandler)

		//galleryKind
		v3NoCache.GET("/galleryKind_list", controller.ResponseGalleryKindListHandler)
		v3NoCache.POST("/galleryKind", controller.CreateGalleryKindHandler)
		v3NoCache.DELETE("/galleryKind", controller.DeleteGalleryKindHandler)
		v3NoCache.PUT("/galleryKind", controller.UpdateGalleryKindHandler)
	}

	v4 := r.Group("/api/keyword")
	{
		v4.POST("/search", controller.ResponseDataAboutSearchKeyword)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSONP(404, gin.H{
			"msg": "404",
		})
	})
	return r
}
