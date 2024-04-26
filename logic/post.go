package logic

import (
	"fmt"
	"go.uber.org/zap"
	"go_blog/dao/mysql"
	"go_blog/dao/redis"
	"go_blog/models"
	"go_blog/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	//雪花算法生成帖子id
	p.PostID = snowflake.GenID()
	//保存到数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return
	}
	return redis.CreatePost(p.PostID, *p.CommunityID)

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
	community, err := mysql.GetCommunityDetail(*post.CommunityID)
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
func GetPostList(page, size int64) ([]*models.ApiPostDetail, error) {
	posts, err := mysql.GetPostList(page, size)
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
		community, err := mysql.GetCommunityDetail(*post.CommunityID) //获取帖子响应的社区信息
		if err != nil {
			zap.L().Error("GetCommunityDetail failed", zap.Error(err),
				zap.Int64("community_id", *post.CommunityID))
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

func GetPostList2(p *models.ParamPostList) ([]*models.ApiPostDetail, error) {
	////如果CommunityID不为默认值，则按社区获取帖子
	var (
		ids []string
		err error
	)
	if p.CommunityID == -1 {
		ids, err = redis.GetPostIDInOrder(p) //从redis中获取排序的帖子ID
	} else {
		ids, err = redis.GetCommunityPostIDInOrder(p)
	}
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDInOrder(p) success, but empty")
		return nil, nil
	}
	posts, err := mysql.GetPostListByID(ids) // 从mysql获取有序的帖子信息
	if err != nil {
		return nil, err
	}
	votes, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}
	dataList := make([]*models.ApiPostDetail, 0, len(posts)) // 通过make创建slice并声明容量以限定slice初始化占用的内存空间
	for idx, post := range posts {                           //获取帖子列表
		user, err := mysql.GetUserByID(post.AuthorID) //获取帖子响应的作者信息
		if err != nil {
			zap.L().Error("GetUserByID failed", zap.Error(err),
				zap.Int64("author_id", post.AuthorID))
			continue //遍历帖子列表，不因为一个帖子的信息获取错误而中断整个循环
		}
		community, err := mysql.GetCommunityDetail(*post.CommunityID) //获取帖子响应的社区信息
		if err != nil {
			zap.L().Error("GetCommunityDetail failed", zap.Error(err),
				zap.Int64("community_id", *post.CommunityID))
			continue
		}
		data := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Votes:           votes[idx],
			CommunityDetail: community,
			Post:            post,
		}
		dataList = append(dataList, data)
	}
	return dataList, nil
}

//确保展示的帖子是最新的或者热度高的
//一天有86400秒，当帖子的票数大于200 那么帖子就存活到下一天，因此一票为86400/200=432分

//为了减少服务器压力，每个帖子仅在有效期内允许投票，到期后将票数持久化存储，未到期的数据保存在redis中

// VotePost 投票分数算法的实现
func VotePost(userID int64, p *models.ParamVote) error {
	return redis.VoteForPost(fmt.Sprintf("%d", userID), p.PostID, float64(*p.Direction))
}
