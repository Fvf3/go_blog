package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

var ErrorUserNotFound = errors.New("用户未登录")

const CtxUIDKey = "userID" //全局常量相较于明文具有可复用性
func getCurrentUser(c *gin.Context) (uid int64, err error) {
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
func getPageInfo(c *gin.Context) (offset, limit int64) {
	//获取分页数
	var err error
	offsetStr := c.Query("offset") //第几页开始
	limitStr := c.Query("limit")   //每页有几条
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 1
	}
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 10
	}
	return
}
