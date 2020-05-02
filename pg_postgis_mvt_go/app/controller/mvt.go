package controller

import (
	"fmt"
	"pg_postgis_mvt/app/service"

	"github.com/kataras/iris/v12"
)

type mvtController struct {
}

// Mvt 矢量切片
var Mvt mvtController

// Export 导出函数
func (ec *mvtController) GetMvt(ctx iris.Context) {
	tableName := ctx.Params().Get("tableName")
	x, _ := ctx.Params().GetInt("x")
	y, _ := ctx.Params().GetInt("y")
	z, _ := ctx.Params().GetInt("z")
	result := service.Mvt.GetMvt(tableName, x, y, z)
	// ctx.Header("Content-Encoding", "deflate")
	ctx.Header("Content-Length", fmt.Sprintf("%d", len([]byte(result))))
	ctx.Header("Content-Type", "application/x-protobuf")
	ctx.Write([]byte(result))
	// ctx.Write([]byte(result))

}
