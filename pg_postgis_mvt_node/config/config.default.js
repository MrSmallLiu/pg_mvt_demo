/* eslint valid-jsdoc: "off" */

'use strict';

/**
 * @param {Egg.EggAppInfo} appInfo app info
 */
module.exports = appInfo => {
  /**
   * built-in config
   * @type {Egg.EggAppConfig}
   **/
  const config = exports = {};
  config.sequelize = {
    dialect: 'postgres', // support: mysql, mariadb, postgres, mssql

    host: '*.*.*.*',
    database: '***',

    port: '9999',
    username: 'postgres',
    password: '********',
    timezone: '+08:00',
    pool: {
      max: 5,
      min: 0,
      idle: 10000,
    },
    logging: console.log,
  };
  config.cors = {
    origin: '*',
    allowMethods: 'GET,HEAD,PUT,POST,DELETE,PATCH',
  };
  // use for cookie sign key, should change to your own and keep security
  config.keys = appInfo.name + '_1588407986000_311';

  // add your middleware config here
  config.middleware = [];

  // add your user config here
  const userConfig = {
    // myAppName: 'egg',
  };

  return {
    ...config,
    ...userConfig,
  };
};
