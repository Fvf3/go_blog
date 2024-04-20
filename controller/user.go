package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"go_blog/logic"
	"go_blog/models"
	"net/http"
)

// 处理注册请求
func SignUpHandler(c *gin.Context) {
	//参数处理
	p := new(models.ParamSignUp)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("Sign up param invalid", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors) //只翻译validator校验错误的内容
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //只有ValidationErrors才有Translate，函数传翻译器, 去除错误中的结构体名称
		})
		return
	}
	//校验请求参数是否符合业务规则在validator中进行
	fmt.Println(p)
	//业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("Sign up process error", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("注册失败，%v", err.Error()),
		})
		return
	}
	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}
func LoginHandler(c *gin.Context) {
	//参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("Login param invalid", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		}
		return
	}
	//逻辑处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("Login process error", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("登录失败，%v", err.Error()),
		})
		return
	}
	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "还需要设置session或者cookie记录登陆状态",
	})
}
