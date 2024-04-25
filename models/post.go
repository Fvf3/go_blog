package models

import "time"

type Post struct { //注意内存对齐
	ID          int64     `json:"id" db:"id"`
	PostID      int64     `json:"post_id" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"` //required会默认参数中该类型的默认值为空，也就是当CommunityID提交了0时，会报错，解决方法是使用指针类型
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type ApiPostDetail struct {
	AuthorName       string             `json:"author_name"`
	*CommunityDetail `json:"community"` //嵌入社区详情  red返回的json会按绑定的属性进行分层
	*Post                               //嵌入帖子详情
}
