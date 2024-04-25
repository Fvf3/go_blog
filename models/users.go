package models

type User struct {
	UserID   int64  `db:"user_id"`  //雪花算法生成的userid
	Username string `db:"username"` // db为sqlx的tag，指定数据库中对应的属性名
	Password string `db:"password"`
	Token    string
}
