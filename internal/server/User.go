package server

import (
	"TSM-Server/internal/serializer"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	LoginTime    string `json:"loginTime"`
	RegisterTime string `json:"registerTime"`
	jwt.StandardClaims
}

func (service *User) Login() serializer.Response {
	err := verify(service.Username, service.Password)
	if err != nil {
		return serializer.Response{
			Msg:   "登陆失败",
			Error: err.Error(),
		}
	}

	user, err := readUser()
	if err != nil {
		return serializer.Response{
			Msg:   "获取用户信息失败",
			Error: err.Error(),
		}
	}

	if comparePasswords(user.Password, service.Password) && user.Username == service.Username {
		token, err := generateToken(user)
		if err != nil {
			return serializer.Response{
				Msg:   "登陆失败",
				Error: err.Error(),
			}
		}

		return serializer.Response{
			Data: token,
			Msg:  "登陆成功",
		}
	}
	return serializer.Response{
		Msg: "登陆失败",
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
	user.Password, _ = hashAndSalt(service.Password)
	user.LoginTime = time.Now().Format("2006-01-02 15:04:05")
	user.RegisterTime = time.Now().Format("2006-01-02 15:04:05")
	file, err := os.OpenFile("./data/user/user.json", os.O_WRONLY|os.O_CREATE, 0666)
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

func (service *User) Logout() serializer.Response {

	return serializer.Response{
		Msg: "登出成功",
	}
}

func verify(username, password string) error {
	if username == "" || password == "" {
		return errors.New("用户名或密码不能为空")
	}
	return nil
}

type Claims struct {
	UserID         interface{}
	Username       string
	StandardClaims jwt.StandardClaims
}

func generateToken(user User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token validity for 24 hours

	claims := jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    "TSM-Server",
		Subject:   user.Username,
	}

	// 使用秘钥签名令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte("TSM-Server")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func hashAndSalt(plaintext string) (pwdHash string, err error) {
	bytePwd := []byte(plaintext)
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	pwdHash = string(hash)
	return pwdHash, nil
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}

func readUser() (User, error) {
	// 读取 ./data/user/user.json 文件
	file, err := os.Open("./data/user/user.json")
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
