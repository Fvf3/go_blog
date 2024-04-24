package logic

import (
	"go_blog/dao/mysql"
	"go_blog/models"
	"go_blog/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	//雪花算法生成帖子id
	p.ID = snowflake.GenID()
	//保存到数据库
	return mysql.CreatePost(p)
}

func GetPost(post_id int64) (data *models.ApiPostDetail, err error) {
	//查询数据，并组合为ApiPostDetail类型
	post, err := mysql.GetPost(post_id)
	if err != nil {
		return
	}
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		return
	}
	community, err := mysql.GetCommunityDetail(post.CommunityID)
	if err != nil {
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		CommunityDetail: community,
		Post:            post,
	}
	return
}
