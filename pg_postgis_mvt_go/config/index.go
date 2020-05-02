package config

import (
	"log"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once

// Viper 读取config文件
// var Viper *viper.Viper

// Config config结构体
type Config struct {
	Port   string
	DBInfo map[string]string
}

// ConfigInfo 配置信息
var ConfigInfo Config

func init() {
	once.Do(func() {
		viper := viper.New()
		// scan the file named config in the root directory
		viper.AddConfigPath("./")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		// read config, if failed, configure by default
		if err := viper.ReadInConfig(); err == nil {
			log.Println("Read config successfully: ", viper.ConfigFileUsed())
			ConfigInfo.Port = viper.GetString("port")
			ConfigInfo.DBInfo = viper.GetStringMapString("database")
		} else {
			log.Printf("Read failed: %s \n", err)
			panic(err)
		}
	})
}

// GetDBConnStr 获取数据库连接字符串
func (config Config) GetDBConnStr() string {
	//TODO: 需要验证没有填写用户名和密码
	host := "host=" + config.DBInfo["host"]
	port := "port=" + config.DBInfo["port"]
	user := "user=" + config.DBInfo["user"]
	dbname := "dbname=" + config.DBInfo["dbname"]
	password := "password=" + config.DBInfo["password"]
	sslmode := "sslmode=" + config.DBInfo["sslmode"]
	str := []string{host, port, user, dbname, password, sslmode}
	return strings.Join(str, " ")
}
