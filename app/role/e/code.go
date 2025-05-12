package e

const (
	SUCCESS      = 200
	ERROR        = 500
	ServiceError = 400

	ErrorDatabase = 40010

	// 认证相关错误
	ErrorAuthCheckTokenFail    = 10001
	ErrorAuthCheckTokenTimeout = 10002
	ErrorAuthTokenInvalid      = 10003

	// 用户相关错误
	ErrorUserNotExist      = 20001
	ErrorUserPasswordWrong = 20002
	ErrorUserAlreadyExist  = 20003

	// 参数错误
	ErrorInvalidParams = 30001
)

var msg = map[int]string{
	SUCCESS:                    "成功",
	ERROR:                      "失败",
	ServiceError:               "服务层错误",
	ErrorAuthCheckTokenFail:    "Token验证失败",
	ErrorAuthCheckTokenTimeout: "Token已过期",
	ErrorAuthTokenInvalid:      "无效的Token",
	ErrorUserNotExist:          "用户不存在",
	ErrorUserPasswordWrong:     "密码错误",
	ErrorUserAlreadyExist:      "用户已存在",
	ErrorInvalidParams:         "参数错误",
}

func GetMsg(code int) string {
	if m, ok := msg[code]; ok {
		return m
	}
	return msg[ERROR]
}
