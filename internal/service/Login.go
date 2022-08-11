package service

import (
	"TSM-Server/utils"
	"strings"
)

type LoginService struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (service *LoginService) Login() bool {

	lines, _ := utils.ReadUserList()
	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], service.Username) && strings.Contains(lines[i], service.Password) {
			return true
		}
	}
	return false

}
