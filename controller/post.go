package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"go_blog/logic"
	"go_blog/models"
	"strconv"
)

const (
	orderTime  = "time"
	orderScore = "score"
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

// GetPostListHandler2 依据前端参数动态获取帖子列表
func GetPostListHandler2(c *gin.Context) {
	//1.参数处理
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: orderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("get post list failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2.从redis查询pid
	//3.依据pid从mysql获取帖子详情
}

// VotePostHandler 为帖子点赞或点踩
func VotePostHandler(c *gin.Context) {
	p := new(models.ParamVote)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("post value bind error", zap.Error(err))
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok { //其他错误
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans))) //参数校验不通过
		return
	}
	userID, err := getCurrentUser(c)
	if err != nil { //获取用户id失败
		zap.L().Error("get current user id failed", zap.Error(err))
		ResponseError(c, CodeAuthEmpty)
		return
	}
	if err := logic.VotePost(userID, p); err != nil { //更新分数和投票记录失败
		zap.L().Error("logic.VotePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
