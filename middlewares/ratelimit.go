package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap) //设置填充速率，总容量
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) == 0 { // 取不到令牌，返回零，此时应当结束对请求的响应
			c.String(http.StatusOK, "rate limit")
			c.Abort()
			return
		}
		c.Next() //取到令牌，继续处理请求
	}
}
