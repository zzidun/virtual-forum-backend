package response

type ResponseCode int64

const (
	CodeSuccess         ResponseCode = 1000
	CodeInvalidParams   ResponseCode = 1001
	CodeUserExist       ResponseCode = 1002
	CodeUserNotExist    ResponseCode = 1003
	CodeInvalidPassword ResponseCode = 1004
	CodeServerBusy      ResponseCode = 1005

	CodeInvalidToken      ResponseCode = 1006
	CodeInvalidAuthFormat ResponseCode = 1007
	CodeNotLogin          ResponseCode = 1008

	CodeInvalidRequestFormat ResponseCode = 2000
	CodeUnknownError         ResponseCode = 9999
)

var msgFlags = map[ResponseCode]string{
	CodeSuccess:         "success",
	CodeInvalidParams:   "请求参数错误",
	CodeUserExist:       "用户名重复",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",

	CodeInvalidToken:      "无效的Token",
	CodeInvalidAuthFormat: "认证格式有误",
	CodeNotLogin:          "未登录",

	CodeInvalidRequestFormat: "请求报文格式错误",
	CodeUnknownError:         "未知错误类型",
}

func (c ResponseCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
