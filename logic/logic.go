package logic

import (
	"go_blog/dao/mysql"
	"go_blog/pkg/snowflake"
)

func SignUp() {
	//查询用户名是否存在
	mysql.QueryUserByUsername()
	//检测密码安全性
	//生成UID
	snowflake.GenID()
	//写入数据库
	mysql.InsertUser()
}
