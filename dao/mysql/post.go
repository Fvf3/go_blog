package mysql

import (
	"github.com/jmoiron/sqlx"
	"go_blog/models"
	"strings"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(post_id,title,content,author_id,community_id) values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.PostID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPost(post_id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post where post_id=?`
	err = db.Get(post, sqlStr, post_id)
	return
}

func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id,create_time from post order by create_time desc  limit ?,?`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

func GetPostListByID(ids []string) (posts []*models.Post, err error) {
	//where指定了查找的pid内容 可以是string切片，而order by find——in——set的参数2指定了排序方式，但参数需要是string
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post where post_id in (?) order by FIND_IN_SET(post_id,?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ",")) //通过In填充sql语句的内容，并返回结果
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query) //通过Rebind将sql语句转为当前数据库格式的语句
	err = db.Select(&posts, query, args...)
	return
}
