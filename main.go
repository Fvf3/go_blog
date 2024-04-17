package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go_blog/dao/mysql"
	"go_blog/dao/redis"
	"go_blog/logger"
	"go_blog/routers"
	"go_blog/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// go web通用开发模板
func main() {
	var configName string
	flag.StringVar(&configName, "configname", "dafault", "config file name") //通过flag获取命令行中指定的配置文件名，传给绑定的变量
	flag.Parse()                                                             //解析命令行参数
	//1.加载配置
	if err := settings.Init(configName); err != nil {
		fmt.Printf("settings init error:%s", err.Error())
		return
	}
	//2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("logger init error:%s", err.Error())
		return
	}
	defer zap.L().Sync() //将缓冲区信息加载到logger中
	zap.L().Debug("logger init success")
	//3.初始化持久化数据库连接
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Printf("mysql init error:%s", err.Error())
		return
	}
	defer mysql.Close()
	//4.初始化缓存数据库连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("redis init error:%s", err.Error())
		return
	}
	defer redis.Close()
	//5.注册路由
	r := routers.Setup()
	//6.启动服务（优雅关机)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: ", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内未处理完的请求处理完再关闭服务，超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
