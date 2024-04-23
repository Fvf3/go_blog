package middlewares

import (
	"github.com/gin-gonic/gin"
	"go_blog/controller"
	"go_blog/pkg/jwt"
	"strings"
)

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
		c.Set(controller.CtxUIDKey, mc.UserID)
		c.Next() //中间件，接下来处理该路由中的其他函数
	}
}
