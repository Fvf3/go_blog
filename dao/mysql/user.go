package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"go_blog/models"
)

const secret = "fvf3"

// CheckUserExist 查询用户名是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username=?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return
	}
	if count > 0 {
		return errors.New("user already exist")
	}
	return
}

// InsertUser 插入用户
func InsertUser(user *models.User) (err error) {
	//数据库中密码不应明文存储，此处加盐
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	password := encrypt(user.Password, secret)
	_, err = db.Exec(sqlStr, user.UserID, user.Username, password)
	return
}

func encrypt(plainText, key string) string {
	h := md5.New()
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum([]byte(plainText)))
}
