package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("用户名已存在")
	ErrorUserNotExist    = errors.New("用户名不存在")
	ErrorPasswordInvalid = errors.New("密码错误")
	ErrorIDInvalid       = errors.New("ID错误")
)
