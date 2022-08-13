package server

import (
	"TSM-Server/internal/api"
	"TSM-Server/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	//登录
	r.POST("login", nil)
	//上传
	r.POST("upload", api.UploadFile)

	v0 := r.Group("api")

	//服务器
	v1 := v0.Group("server")
	{
		//服务器信息
		v1.POST("Info", api.GetServerInfo)
		//时间
		v1.POST("setTime", api.SetTime)
		//开服、关服、重启
		v1.POST("action", api.ServerAction)
	}

	//模组
	v2 := v0.Group("mod")
	{
		//模组信息
		v2.POST("Info", api.GetModInfo)
		//启用、禁用
		v2.POST("action", api.ModAction)
		//删除
		v2.POST("del", api.DelMod)
	}

	//玩家
	v3 := v0.Group("player")
	{
		//玩家信息
		v3.POST("Info", api.GetPlayerInfo)
		//踢出
		v3.POST("kick", api.KicPlayer)
		//加入黑名单
		v3.POST("block", api.BlockPlayer)
		//从黑名单删除
		v3.POST("del", api.DelPlayer)
	}

	return r
}
