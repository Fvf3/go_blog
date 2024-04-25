package logic

import (
	"go.uber.org/zap"
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

// GetPostList 获取帖子详情的列表
func GetPostList(offset, limit int64) ([]*models.ApiPostDetail, error) {
	posts, err := mysql.GetPostList(offset, limit)
	if err != nil {
		zap.L().Error("GetPostList failed", zap.Error(err))
		return nil, err
	}
	dataList := make([]*models.ApiPostDetail, 0, len(posts)) // 通过make创建slice并声明容量以限定slice初始化占用的内存空间
	for _, post := range posts {                             //获取帖子列表
		user, err := mysql.GetUserByID(post.AuthorID) //获取帖子响应的作者信息
		if err != nil {
			zap.L().Error("GetUserByID failed", zap.Error(err),
				zap.Int64("author_id", post.AuthorID))
			continue //遍历帖子列表，不因为一个帖子的信息获取错误而中断整个循环
		}
		community, err := mysql.GetCommunityDetail(post.CommunityID) //获取帖子响应的社区信息
		if err != nil {
			zap.L().Error("GetCommunityDetail failed", zap.Error(err),
				zap.Int64("community_id", post.CommunityID))
			continue
		}
		data := &models.ApiPostDetail{
			AuthorName:      user.Username,
			CommunityDetail: community,
			Post:            post,
		}
		dataList = append(dataList, data)
	}
	return dataList, nil
}
