package serializer

import "log"

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

func HandleErr(err error, msg string) Response {
	log.Printf("错误: %v", err)
	return Response{
		Status: 500,
		Error:  err.Error(),
		Msg:    msg,
	}
}
