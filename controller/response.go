package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseError 当发生错误时返回的响应
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  code.Msg(),
	})
}

// ResponseSuccess 正常的响应
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  CodeSuccess.Msg(),
		"data": data,
	})
}

// ResponseErrorWithMsg 发生错误时，返回带有消息的响应
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
