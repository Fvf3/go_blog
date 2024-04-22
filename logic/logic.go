package logic

import (
	"go_blog/dao/mysql"
	"go_blog/models"
	"go_blog/pkg/jwt"
	"go_blog/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) error {
	//查询用户名是否存在
	exist, err := mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	if exist {
		return mysql.ErrorUserExist
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
func Login(p *models.ParamLogin) (token string, err error) {
	exist, err := mysql.CheckUserExist(p.Username)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", mysql.ErrorUserNotExist
	}
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err = mysql.CheckPasswordCorrect(user); err != nil {
		return "", err
	}
	//生成JWT token
	return jwt.GenToken(user.UserID, user.Username)

}
