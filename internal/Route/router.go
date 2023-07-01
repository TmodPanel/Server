package Route

import (
	"TSM-Server/internal/api"
	"TSM-Server/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	v1Api := r.Group("/api/v1")

	v1Api.POST("/login", api.Login)       // 登录
	v1Api.POST("/register", api.Register) // 注册

	authGroup := v1Api.Group("/auth")
	authGroup.Use(middleware.AuthMiddleware()) // 添加身份验证中间件
	{
		gameApi := v1Api.Group("/game")
		{
			gameApi.GET("/:id", api.GetGameInfo)          // 获取游戏信息
			gameApi.PUT("/:id/time", api.SetTime)         // 设置游戏时间
			gameApi.POST("/:id/action", api.ServerAction) // 执行游戏操作
		}

		modApi := v1Api.Group("/mod")
		{
			modApi.GET("/:id", api.GetModInfo)        // 获取模块信息
			modApi.POST("/:id/action", api.ModAction) // 执行模块操作
			modApi.DELETE("/:id", api.DelMod)         // 删除模块
			modApi.GET("/", api.GetModList)           // 获取模块列表
		}

		playerApi := v1Api.Group("/player")
		{
			playerApi.GET("/:id", api.GetPlayerInfo)      // 获取玩家信息
			playerApi.POST("/:id/kick", api.KickPlayer)   // 踢出玩家
			playerApi.POST("/:id/block", api.BlockPlayer) // 封禁玩家
			playerApi.DELETE("/:id", api.DelPlayer)       // 删除玩家
		}

		configApi := v1Api.Group("/config")
		{
			configApi.GET("/:id", api.GetSchemesInfo)     // 获取配置信息
			configApi.DELETE("/:id", api.DelScheme)       // 删除配置
			configApi.POST("/", api.AddScheme)            // 添加配置
			configApi.PUT("/:id", api.UpdateScheme)       // 更新配置
			configApi.POST("/:id/reset", api.ResetScheme) // 重置配置
		}

		serverApi := v1Api.Group("/server")
		{
			serverApi.GET("/:id", api.GetServerInfo) // 获取服务器信息
		}

		fileApi := v1Api.Group("/file")
		{
			fileApi.GET("/", api.GetFileList)     // 获取文件列表
			fileApi.DELETE("/:id", api.DelFile)   // 删除文件
			fileApi.POST("/", api.UploadFile)     // 上传文件
			fileApi.GET("/:id", api.DownloadFile) // 下载文件
		}
	}

	miscApi := v1Api.Group("/misc")
	{
		miscApi.GET("/version", func(context *gin.Context) { // 获取版本信息
			context.JSON(200, gin.H{"version": "1.0.0"})
		})
		miscApi.GET("/ping", func(context *gin.Context) { // ping
			context.JSON(200, gin.H{"message": "pong"})
		})
	}
	return r
}
