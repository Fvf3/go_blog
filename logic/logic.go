package logic

import (
	"errors"
	"go_blog/dao/mysql"
	"go_blog/models"
	"go_blog/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) error {
	//查询用户名是否存在
	exist, err := mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("用户名已存在")
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
	return mysql.InsertUser(user)
}
func Login(p *models.ParamLogin) error {
	exist, err := mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("用户名不存在")
	}
	return mysql.CheckPasswordCorrect(p.Username, p.Password)
}
