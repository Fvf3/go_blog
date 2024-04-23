package mysql

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"go_blog/models"
)

func GetCommunityList() (list []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err = db.Select(&list, sqlStr); err != nil {
		if errors.Is(err, sql.ErrNoRows) { //没有社区,属于正常现象，无需返回错误
			zap.L().Warn("community list empty")
			err = nil
			return
		}
	}
	return
}

func GetCommunityDetail(id int64) (communityDetail *models.CommunityDetail, err error) {
	sqlStr := "select community_id,community_name,introduction,create_time from community where community_id = ?"
	communityDetail = new(models.CommunityDetail) //当需要传入指针以获取值时，应当在方法内部声明指针，在函数声明中声明的指针会被赋予nil，因此无法用以获取值
	if err = db.Get(communityDetail, sqlStr, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) { //社区不存在
			err = ErrorIDInvalid
		}
		return
	}
	return
}
