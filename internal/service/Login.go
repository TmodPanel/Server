package service

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type LoginService struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (service *LoginService) Login() {
	userList, _ := ReadUserList("./config/userList.txt")
	flag := false
	for _, v := range userList {
		if v.Username == service.Username && v.Password == service.Password {
			flag = true
		}
	}
	if !flag {
		fmt.Println("用户名或密码错误")
	} else {
		fmt.Println("登陆成功")
	}

}

func ReadUserList(path string) ([]LoginService, error) {
	var userList []LoginService
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	buf := bufio.NewScanner(file)
	for {
		if !buf.Scan() {
			break
		}
		line := buf.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		slice := strings.Split(line, "=")
		user := LoginService{
			Username: strings.Split(slice[1], " ")[0],
			Password: strings.Split(slice[2], " ")[0],
		}
		userList = append(userList, user)
	}
	return userList, err
}
