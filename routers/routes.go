package routers

import (
	"github.com/gin-gonic/gin"
	"go_blog/controller"
	"go_blog/logger"
	"go_blog/middlewares"
	"net/http"
	"time"
)

func Setup(mode string) *gin.Engine {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1)) //通过在指定的路由方法添加认证中间件，校验是否登陆
	v1 := r.Group("/api/v1")
	v1.POST("/signup", controller.SignUpHandler) //注册
	v1.POST("/login", controller.LoginHandler)   //登陆以获取token
	v1.Use(middlewares.JWTAuthMiddleware())      //应用jwt token中间件
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostHandler)
		v1.GET("/posts", controller.GetPostListHandler)
		v1.POST("/vote", controller.VotePostHandler)
		v1.GET("/posts2", controller.GetPostList2Handler)
	}
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	return r
}
