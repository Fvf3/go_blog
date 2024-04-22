package controller

// ResCode 定义了返回请求的状态码
type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeAuthEmpty
	CodeAuthInvalid
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求的参数无效",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "密码错误",
	CodeServerBusy:      "服务器繁忙", //并不向用户返回真正的错误信息
	CodeAuthEmpty:       "token为空",
	CodeAuthInvalid:     "token异常",
}

// Msg 获取状态码对应的消息
func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
