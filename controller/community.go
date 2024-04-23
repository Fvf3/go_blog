package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_blog/logic"
	"strconv"
)

// CommunityHandler 以列表形式返回所有社区的名称与ID
func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 获取社区详情
func CommunityDetailHandler(c *gin.Context) {
	//获取社区ID
	idStr := c.Param("id")
	id, ok := strconv.ParseInt(idStr, 10, 64)
	if ok != nil { //参数无效
		ResponseError(c, CodeInvalidParam)
		return
	}
	//获取社区详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil { //查询失败
		zap.L().Error("logic.GetCommunityDetail failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
