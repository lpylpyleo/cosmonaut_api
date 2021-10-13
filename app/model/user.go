package model

// 注册请求参数，用于前后端交互参数格式约定
type UserApiSignUpReq struct {
	Email    string `v:"required|email#账号不能为空|账号应为合法email"`
	Password string `v:"required#请输入确认密码"`
	Nickname string `v:"length:6,20"`
}

// 登录请求参数，用于前后端交互参数格式约定
type UserApiSignInReq struct {
	Email    string `v:"required#账号不能为空"`
	Password string `v:"required#密码不能为空"`
}

// 账号唯一性检测请求参数，用于前后端交互参数格式约定
type UserApiCheckEmailReq struct {
	Email string `v:"required#账号不能为空"`
}

// 昵称唯一性检测请求参数，用于前后端交互参数格式约定
type UserApiCheckNickNameReq struct {
	Nickname string `v:"required#昵称不能为空"`
}

// 注册输入参数
type UserServiceSignUpReq struct {
	Id       string
	Email    string
	Password string
}
