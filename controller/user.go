package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"go_blog/dao/mysql"
	"go_blog/logic"
	"go_blog/models"
)

// 处理注册请求
func SignUpHandler(c *gin.Context) {
	//参数处理
	p := new(models.ParamSignUp)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("Sign up param invalid", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors) //只翻译validator校验错误的内容
		if !ok {
			ResponseError(c, CodeServerBusy) //其他错误
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans))) //校验出错
		return
	}
	//校验请求参数是否符合业务规则在validator中进行
	fmt.Println(p)
	//业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("Sign up process error", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist) //用户已存在
		} else {
			ResponseError(c, CodeServerBusy) //数据库操作出错
		}
		return
	}
	ResponseSuccess(c, nil)
}

// 处理登陆请求，返回token
func LoginHandler(c *gin.Context) {
	//参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("Login param invalid", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans))) //校验出错
		} else {
			ResponseError(c, CodeServerBusy) // 其他错误
		}
		return
	}
	//逻辑处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("Login process error", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist) //用户不存在
		} else if errors.Is(err, mysql.ErrPasswordInvalid) {
			ResponseError(c, CodeInvalidPassword) //密码错误
		} else {
			ResponseError(c, CodeServerBusy) //数据库查询出错
		}
		return
	}
	ResponseSuccess(c, token)
}
