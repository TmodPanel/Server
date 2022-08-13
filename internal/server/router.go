package server

import (
	"TSM-Server/internal/service"
	"TSM-Server/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	//登录
	r.POST("login", nil)
	//上传
	r.POST("upload", service.UploadFile)

	api := r.Group("api")

	//服务器
	v1 := api.Group("server")
	{
		//服务器信息
		v1.POST("Info", service.GetServerInfo)
		//时间
		v1.POST("setTime", service.SetTime)
		//开服、关服、重启
		v1.POST("action", service.ServerAction)
	}

	//模组
	v2 := api.Group("mod")
	{
		//模组信息
		v2.POST("Info", service.GetModInfo)
		//启用、禁用
		v2.POST("action", service.ModAction)
		//删除
		//v2.POST("del", service.DelMod)
	}

	//玩家
	v3 := api.Group("player")
	{
		//玩家信息
		v3.POST("Info", service.GetPlayerInfo)
		//踢出
		v3.POST("kick", service.KicPlayer)
		//加入黑名单
		v2.POST("block", service.BlockPlayer)
		//从黑名单删除
		v2.POST("del", service.DelPlayer)
	}

	return r
}
