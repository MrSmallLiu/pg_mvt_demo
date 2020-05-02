'use strict';

const Service = require('egg').Service;
const util = require('../../util')
class MvtService extends Service {
    async getMvt(tableName, x, y, z) {
        const result = {
            data: null,
            code: 1,
            msg: "查询成功"
        }
        x = parseInt(x);
        y = parseInt(y);
        z = parseInt(z);
        let [xmin, ymin] = util.xyz2lonlat(x, y, z);
        let [xmax, ymax] = util.xyz2lonlat(x + 1, y + 1, z);
        const randomMap = {
            4: 0.1,
            5: 0.2,
            6: 0.3,
            7: 0.45,
            8: 0.6,
            9: 0.75,
        };
        let randomFilter = ''
        if (z <= 4) {
            randomFilter = ' and random < 0.1';
        } else if (z >= 10) {
            randomFilter = '';
        } else {
            randomFilter = ` and random < ${randomMap[z]}`;
        }
        //组织SQL
        let sql1 = ` SELECT
        ST_AsMVT ( P,'points', 4096, 'geom' ) AS "mvt" FROM	(SELECT 
          ST_AsMVTGeom (ST_Transform (geom, 3857 ),ST_Transform (	ST_MakeEnvelope
          ( ${xmin},${ymin}, ${xmax},${ymax}, 4326 ),3857),
          4096,	64,TRUE ) geom FROM "${tableName}" where sjzt='1' and geom && ST_MakeEnvelope
          ( ${xmin},${ymin}, ${xmax},${ymax}, 4326 ) ${randomFilter} ) AS P `;
        let sql2 = ` SELECT
        ST_AsMVT ( P,'line', 4096, 'geom' ) AS "mvt" FROM	(SELECT 
          ST_AsMVTGeom (ST_Transform (geom, 3857 ),	ST_Transform (ST_MakeEnvelope
          ( ${xmin},${ymin}, ${xmax},${ymax}, 4326 ),3857),
          4096,	64,TRUE ) geom FROM "data_10001001538"  ) AS P `
          let sql3 = ` SELECT
        ST_AsMVT ( P,'polygon', 4096, 'geom' ) AS "mvt" FROM	(SELECT 
          ST_AsMVTGeom (ST_Transform (st_simplify(geom,0.0), 3857 ),	ST_Transform (ST_MakeEnvelope
          ( ${xmin},${ymin}, ${xmax},${ymax}, 4326 ),3857),
          4096,	64,TRUE ) geom FROM "data_10001001532"  ) AS P `
        const res = await this.ctx.model.query(`select (${sql1})||(${sql2})||(${sql3}) as mvt`).catch(err => {
            console.error(err)
            result.code = 0;
            result.data = err;
            result.msg = "查询错误"
        })
        if (!res) return result;
        if (res[0] && res[0][0] && res[0][0].mvt && res[0][0].mvt.length > 0) {
            result.data = res[0][0].mvt;
        } else {
            result.code = 0;
        }
        return result;
    }
}

module.exports = MvtService;
