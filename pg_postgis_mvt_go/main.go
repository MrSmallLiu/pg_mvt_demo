package main

import (
	"pg_postgis_mvt/app/router"
	"pg_postgis_mvt/config"
	"pg_postgis_mvt/app/middleware"
	"github.com/kataras/iris/v12"
)

func main() {
	irisApp := iris.New()
	irisApp.AllowMethods(iris.MethodOptions)
	irisApp.Use(middleware.CORS)
	router.Router(irisApp)
	configInfo := config.ConfigInfo
	irisApp.Listen(configInfo.Port)
}
