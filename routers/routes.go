package routers

import (
	"github.com/gin-gonic/gin"
	"go_blog/controller"
	"go_blog/logger"
	"go_blog/pkg/jwt"
	"net/http"
	"strings"
)

func Setup(mode string) *gin.Engine {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) //通过在指定的路由方法添加认证中间件，校验是否登陆
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler) //登陆以获取token
	r.GET("/hello", JWTAuthMiddleware(), func(c *gin.Context) {
		//仅当用户通过了认证中间件才能进行到该函数
		c.String(http.StatusOK, "hello")
	})
	return r
}

// JWTAuthMiddleware 基于JWT token的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	//用户以beartoken 将token放在请求头
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeAuthEmpty)
			c.Abort() //提前终止请求处理
			return
		}
		parts := strings.SplitN(authHeader, " ", 2) //0: "Bearer" 1: Token
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeAuthInvalid)
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeAuthInvalid)
			c.Abort()
			return
		}
		//将当前请求的用户信息保存到请求的上下文，用于后续路由函数处理
		c.Set("username", mc.UserID)
		c.Next() //中间件，接下来处理该路由中的其他函数
	}
}
