package service

import (
	"fmt"
	"pg_postgis_mvt/database"
	"pg_postgis_mvt/util"
)

type mvtService struct {
}

// Mvt mvt service
var Mvt mvtService

func (*mvtService) GetMvt(tableName string, x int, y int, z int) string {
	// var random string
	randomFilter := ""
	randomMap := map[int]float32{
		4: 0.1,
		5: 0.2,
		6: 0.3,
		7: 0.45,
		8: 0.6,
		9: 0.75,
	}
	if z <= 4 {
		// tableName = 'jcb_cd_4';
		randomFilter = " and random < 0.1"
	} else if z >= 10 {
		randomFilter = ""
	} else {
		// tableName = `jcb_cd_${z}`;
		randomFilter = fmt.Sprintf(" and random < %f", randomMap[z])
	}

	xyMin := util.XYZ2lonlat(x, y, z)
	xyMax := util.XYZ2lonlat(x+1, y+1, z)

	//组织SQL
	sql1 := fmt.Sprintf(`select ST_AsMVT ( P,'points', 4096, 'geom' ) AS "mvt" FROM	(SELECT
	  ST_AsMVTGeom (ST_Transform (geom, 3857 ),ST_Transform (	ST_MakeEnvelope
	  ( %f,%f, %f,%f, 4326 ),3857),
	  4096,	64,TRUE ) geom FROM %s where sjzt='1' and geom && ST_MakeEnvelope
	  ( %f,%f, %f,%f, 4326 ) %s ) AS P`, xyMin[0], xyMin[1], xyMax[0], xyMax[1], tableName, xyMin[0], xyMin[1], xyMax[0], xyMax[1], randomFilter)
	// let sql2 = ` SELECT
	// ST_AsMVT ( P,'line', 4096, 'geom' ) AS "mvt" FROM	(SELECT
	//   ST_AsMVTGeom (ST_Transform (geom, 3857 ),	ST_Transform (ST_MakeEnvelope
	//   ( ${xmin},${ymin}, ${xmax},${ymax}, 4326 ),3857),
	//   4096,	64,TRUE ) geom FROM "data_10001001538"  ) AS P `
	//   let sql3 = ` SELECT
	// ST_AsMVT ( P,'polygon', 4096, 'geom' ) AS "mvt" FROM	(SELECT
	//   ST_AsMVTGeom (ST_Transform (st_simplify(geom,0.02), 3857 ),	ST_Transform (ST_MakeEnvelope
	//   ( ${xmin},${ymin}, ${xmax},${ymax}, 4326 ),3857),
	//   4096,	64,TRUE ) geom FROM "data_10001001532"  ) AS P `
	res, _ := database.DB.QueryString(sql1)
	return (res[0])["mvt"]
}
