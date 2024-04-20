package models

// 请求参数的结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`                     // 必填
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` //必须等于Password
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
