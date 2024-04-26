package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 请求参数的结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`                     // 必填
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` //必须等于Password
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVote struct {
	PostID    string `json:"post_id" binding:"required"`                // 帖子id
	Direction *int8  `json:"direction" binding:"required,oneof=-1 0 1"` //赞(1) 踩(-1) 取消(0)
}

type ParamPostList struct {
	Page        int64  `form:"page"`
	Size        int64  `form:"size"`
	Order       string `form:"order"`
	CommunityID int64  `form:"community_id"`
}
