'use strict';

const Controller = require('egg').Controller;

class MvtController extends Controller {
    async getMvt() {
        const { ctx } = this
        const { tableName, x, y, z } = ctx.params
        const result = await ctx.service.mvt.getMvt(tableName, x, y, z)
        if (result.code === 0) {
            ctx.status = 404;
            return;
        }
        ctx.set('Content-Type', 'application/x-protobuf');
        ctx.body = result.data;
    }
}

module.exports = MvtController;
