package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var ErrorUserNotFound = errors.New("用户未登录")

const CtxUIDKey = "userID" //全局常量相较于明文具有可复用性
func GetCurrentUser(c *gin.Context) (uid int64, err error) {
	userID, ok := c.Get(CtxUIDKey)
	if !ok { //用户信息参数为空
		err = ErrorUserNotFound
		return
	}
	uid, ok = userID.(int64) //gin.Get 返回any类型，需通过断言转换
	if !ok {                 //用户信息无法解析
		err = ErrorUserNotFound
		return
	}
	return
}
