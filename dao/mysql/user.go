package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"go_blog/models"
)

const secret = "fvf3"

// CheckUserExist 查询用户名是否存在
func CheckUserExist(username string) (bool, error) {
	sqlStr := `select count(user_id) from user where username=?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	return count > 0, nil
}

// InsertUser 插入用户
func InsertUser(user *models.User) (err error) {
	//数据库中密码不应明文存储，此处加盐
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	password := encrypt(user.Password, secret)
	_, err = db.Exec(sqlStr, user.UserID, user.Username, password)
	return
}
func CheckPasswordCorrect(username, password string) error {
	newPassword := encrypt(password, secret)
	var oldPassword string
	sqlStr := `select password from user where username=?`
	if err := db.Get(&oldPassword, sqlStr, username); err != nil {
		return err
	}
	if newPassword != oldPassword {
		return errors.New("密码错误")
	}
	return nil
}
func encrypt(plainText, key string) string {
	h := md5.New()
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum([]byte(plainText)))
}
