package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_blog/logic"
	"go_blog/models"
	"strconv"
)

func CreatePostHandler(c *gin.Context) {
	//获取参数
	p := new(models.Post)
	if ok := c.ShouldBind(p); ok != nil {
		zap.L().Error("post value bind error", zap.Error(ok))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//获取用户ID
	AuthorID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeAuthEmpty) //用户ID为空，需要登录
		return
	}
	p.AuthorID = AuthorID
	//逻辑处理
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func GetPostHandler(c *gin.Context) {
	//获取帖子ID
	pid := c.Param("id")
	post_id, ok := strconv.ParseInt(pid, 10, 64)
	if ok != nil {
		zap.L().Error("post id parse error", zap.Error(ok))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//取出帖子数据
	data, err := logic.GetPost(post_id)
	if err != nil {
		zap.L().Error("logic.GetPost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
func GetPostListHandler(c *gin.Context) {
	offset, limit := getPageInfo(c)
	data, err := logic.GetPostList(offset, limit)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
