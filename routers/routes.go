package routers

import (
	"github.com/gin-gonic/gin"
	"go_blog/controller"
	"go_blog/logger"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.POST("/signup", controller.SignUpHandler)
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	return r
}
