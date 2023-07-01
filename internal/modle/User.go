package modle

import (
	"TSM-Server/internal/serializer"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	LoginTime    string `json:"loginTime"`
	RegisterTime string `json:"registerTime"`
}

func (service *User) Login() serializer.Response {
	err := verify(service.Username, service.Password)
	if err != nil {
		return serializer.Response{
			Msg:   "登陆失败",
			Error: err.Error(),
		}
	}

	user, err := ReadUser()
	if err != nil {
		return serializer.Response{
			Msg:   "获取用户信息失败",
			Error: err.Error(),
		}
	}

	if ComparePasswords(user.Password, service.Password) && user.Username == service.Username {
		return serializer.Response{
			Data: user,
			Msg:  "登陆成功",
		}
	} else {
		return serializer.Response{
			Msg: "登陆失败",
		}
	}

}

func (service *User) Register() serializer.Response {
	err := verify(service.Username, service.Password)
	if err != nil {
		return serializer.Response{
			Msg:   "注册失败",
			Error: err.Error(),
		}
	}
	var user User
	user.Username = service.Username
	user.Password, _ = HashAndSalt(service.Password)
	user.LoginTime = time.Now().Format("2006-01-02 15:04:05")
	user.RegisterTime = time.Now().Format("2006-01-02 15:04:05")
	file, err := os.OpenFile("./data/User/user.json", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return serializer.Response{
			Msg:   "注册失败",
			Error: err.Error(),
		}
	}

	defer file.Close()
	buf, _ := json.MarshalIndent(user, "", "\t")
	if _, err := file.Write(buf); err != nil {
		return serializer.Response{
			Msg:   "注册失败",
			Error: err.Error(),
		}
	} else {
		return serializer.Response{
			Msg: "注册成功",
		}
	}

}

func verify(username, password string) error {
	if username == "" || password == "" {
		return errors.New("用户名或密码不能为空")
	}
	return nil
}

func HashAndSalt(plaintext string) (pwdHash string, err error) {
	bytePwd := []byte(plaintext)
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	pwdHash = string(hash)
	return pwdHash, nil
}

func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}

func ReadUser() (User, error) {
	// 读取 ./data/User/user.json 文件
	file, err := os.Open("./data/User/user.json")
	if err != nil {
		return User{}, err
	}
	defer file.Close()
	buf := new(bytes.Buffer)
	io.Copy(buf, file)
	var user User
	err = json.Unmarshal(buf.Bytes(), &user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
