package logic

import (
	"go_blog/dao/mysql"
	"go_blog/models"
	"go_blog/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	//查询用户名是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	//生成UID
	userID := snowflake.GenID()
	//User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//写入数据库
	if err := mysql.InsertUser(user); err != nil {
		return err
	}
	return
}
