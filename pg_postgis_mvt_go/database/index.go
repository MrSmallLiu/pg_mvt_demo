package database

import (
	// "database/sql"
	"log"
	"pg_postgis_mvt/config"
	"sync"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var once sync.Once

// DB 数据库连接实例
var DB *xorm.Engine

// var DB *sql.DB

func init() {
	once.Do(func() {
		var err error
		configInfo := config.ConfigInfo
		dbConncStr := configInfo.GetDBConnStr()
		//"postgres://postgres:geoc_sport@172.16.111.6:9999/tycdcth_test?sslmode=disable"
		DB, err = xorm.NewEngine(configInfo.DBInfo["driver"], dbConncStr)
		// DB, err = sql.Open("postgres", dbConncStr)
		DB.ShowSQL(true)
		if err != nil {
			log.Panicf("数据库连接错误: %s", err)
		}
		// err = DB.Ping()
		// if err != nil {
		// 	log.Panicf("数据库ping错误: %s", err)
		// }
	})
}
