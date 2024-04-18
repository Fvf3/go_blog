package controller

import (
	"github.com/gin-gonic/gin"
	"go_blog/logic"
	"net/http"
)

// 处理注册请求
func SignUpHandler(c *gin.Context) {
	//参数处理
	//业务处理
	logic.SignUp()
	//返回响应
	c.JSON(http.StatusOK, "ok")
}
