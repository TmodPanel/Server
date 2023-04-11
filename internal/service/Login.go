package service

import (
	"TSM-Server/internal/serializer"
	"TSM-Server/utils"
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"strings"
)

type LoginService struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (service *LoginService) Login() serializer.Response {
	m := md5.New()
	io.WriteString(m, service.Password)
	password := hex.EncodeToString(m.Sum(nil))
	log.Println(password)
	lines, err := utils.ReadUserList()
	if err != nil {
		return serializer.Response{
			Msg:   "错误",
			Error: "没有找到密码本",
		}
	}
	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], service.Username) && strings.Contains(lines[i], password) {
			return serializer.Response{
				Msg: "登录成功",
			}
		}
	}
	return serializer.Response{
		Msg:   "登录失败",
		Error: "账号或密码错误",
	}
}
