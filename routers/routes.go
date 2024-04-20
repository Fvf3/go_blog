package routers

import (
	"github.com/gin-gonic/gin"
	"go_blog/controller"
	"go_blog/logger"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	return r
}
