package logic

import (
	"go_blog/dao/mysql"
	"go_blog/models"
)

// GetCommunityList 查询数据库，返回社区列表
func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

// GetCommunityDetail 查询数据库，返回社区详情
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetail(id)
}
