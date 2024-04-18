package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig) //全局的配置文件变量，通过反序列化viper读到的配置得来
type AppConfig struct {
	//配置文件中name等属性需与log同级
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init(filename string) (err error) {
	viper.SetConfigFile(filename) //在命令行传入的参数中找配置文件
	//viper.SetConfigFile("config.yaml") //配置文件名
	//viper.AddConfigPath(".")   //配置查找配置文件的路径
	err = viper.ReadInConfig() // 读取配置文件
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err) //没读到
		return
	}
	if err = viper.Unmarshal(Conf); err != nil { // 反序列化，传给Conf,Conf是new出来的，因此是指针
		fmt.Printf("viper.Unmarshal() failed, err:%v\n", err)
	}
	fmt.Println(Conf.StartTime, Conf.MachineID, Conf.Name, Conf.Port, Conf.Mode)
	viper.WatchConfig()                            //启用支持热加载
	viper.OnConfigChange(func(in fsnotify.Event) { //回调，当配置文件发生修改时触发
		fmt.Println("Config file changed:", in.Name)
		if err = viper.Unmarshal(Conf); err != nil { //配置文件更新了，此时也需要更新全局的配置变量
			fmt.Printf("viper.Unmarshal() failed, err:%v\n", err)
		}
	})
	return
}
