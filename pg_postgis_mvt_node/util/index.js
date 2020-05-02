'use strict';
module.exports = {
  xyz2lonlat(x, y, z) {
    // let n = Math.pow(2, z);
    // let m = Math.pow(2, z+1);
    // let lon_deg = x * 360.0 / m -180;
    // // let lat_rad = Math.atan(Math.sinh(Math.PI * (1 - (2 * y) / n)));
    // let lat_deg =90- y * 180.0/n - 180.0/n;
    // return [lon_deg, lat_deg];
    const n = Math.pow(2, z);
    const lon_deg = (x / n) * 360.0 - 180.0;
    const lat_rad = Math.atan(Math.sinh(Math.PI * (1 - (2 * y) / n)));
    const lat_deg = (180 * lat_rad) / Math.PI;
    return [ lon_deg, lat_deg ];
  },
};
