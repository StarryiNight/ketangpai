package controller

type MyCode int64

const (
	CodeSuccess         MyCode = 1000
	CodeInvalidParams   MyCode = 1001
	CodeUserExist       MyCode = 1002
	CodeUserNotExist    MyCode = 1003
	CodeInvalidPassword MyCode = 1004
	CodeServerBusy      MyCode = 1005

	CodeInvalidToken      MyCode = 1006
	CodeInvalidAuthFormat MyCode = 1007
	CodeNotLogin          MyCode = 1008
	CodeRequestFileError  MyCode = 1009
	CodeRedisSaveFiled    MyCode = 1010
	CodeFileSeekFailed    MyCode = 1011
	CodeLimitedAuthority  MyCode = 1012
)

var msgFlags = map[MyCode]string{
	CodeSuccess:          "success",
	CodeInvalidParams:    "请求参数错误",
	CodeUserExist:        "用户名重复",
	CodeUserNotExist:     "用户不存在",
	CodeInvalidPassword:  "用户名或密码错误",
	CodeServerBusy:       "服务繁忙",
	CodeLimitedAuthority: "用户所在组不能对此操作",

	CodeInvalidToken:      "无效的Token",
	CodeInvalidAuthFormat: "认证格式有误",
	CodeNotLogin:          "未登录",
	CodeRequestFileError:  "获取上传文件错误",
	CodeRedisSaveFiled:    "redis保存失败",
	CodeFileSeekFailed:    "seek长度和文件大小不一致",
}

func (c MyCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
