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

	authGroup := v1Api.Group("/")
	authGroup.Use(middleware.AuthMiddleware()) // 添加身份验证中间件
	{
		instanceApi := authGroup.Group("/instance")
		{
			instanceApi.POST("/", api.AddInstance)        // 添加游戏配置
			instanceApi.GET("/", api.GetInstanceInfo)     // 获取游戏信息
			instanceApi.POST("/timer", api.SetGameTime)   // 设置游戏时间
			instanceApi.POST("/restart", api.RestartGame) // 重启游戏
			instanceApi.POST("/start", api.StartGame)     // 启动游戏
			instanceApi.POST("/stop", api.StopGame)       // 关闭游戏
			instanceApi.DELETE("/", api.DelInstance)      // 删除游戏配置
		}

		modApi := authGroup.Group("/mod")
		{
			modApi.GET("/", api.GetModList)            // 获取模组列表
			modApi.PUT("/:id/enable", api.EnableMod)   // 启用模组
			modApi.PUT("/:id/disable", api.DisableMod) // 禁用模组
			modApi.DELETE("/:id", api.DelMod)          // 删除模组
		}

		playerApi := authGroup.Group("/player")
		{
			playerApi.GET("/", api.GetPlayerInfo)         // 获取玩家列表
			playerApi.POST("/:id/kick", api.KickPlayer)   // 踢出玩家
			playerApi.POST("/:id/block", api.BlockPlayer) // 封禁玩家
			//playerApi.DELETE("/:id", api.DelPlayer)       // 删除玩家
		}

		schemeApi := authGroup.Group("/scheme")
		{
			schemeApi.GET("/", api.GetSchemesInfo) // 获取配置方案
			schemeApi.DELETE("/", api.DelScheme)   // 删除配置
			schemeApi.PUT("/", api.UpdateScheme)   // 更新配置
		}

		panelApi := authGroup.Group("/panel")
		{
			panelApi.GET("/info", api.GetPanelInfo) // 获取面板信息
		}

		fileApi := authGroup.Group("/file")
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
