package serializer

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// HandleErr 处理错误并返回序列化器
func HandleErr(err error, msg string) Response {
	logError(err)
	return Response{
		Status: 500,
		Error:  err.Error(),
		Msg:    msg,
	}
}

// logError 记录错误日志
func logError(err error) {
}
