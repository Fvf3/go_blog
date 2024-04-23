package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

// 秘密
var mySecret = []byte("Don`t tell anybody")

// MyClaims 封装了jwt和用户id用户名
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成jwt
func GenToken(userID int64, username string) (string, error) {
	c := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("auth.jwt_expire")) * time.Minute).Unix(), //过期时间
			Issuer:    "go_frame",                                                                          //签发人
		},
	}
	//指定使用的签名方法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//指定用于签名的秘密，得到字符串格式的token
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) { //解析token，第三个参数为回调函数
		return mySecret, err
	})
	if err != nil { //解析出错
		return nil, err
	}
	if token.Valid { //token有效
		return mc, nil
	}
	return nil, errors.New("token is invalid") //token无效
}
