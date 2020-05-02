package util

import (
	"math"
)

// MapStringDefault MapString取默认值 k: key  d: default
func MapStringDefault(m map[string]string, k string, d string) string {
	result, ok := m[k]
	if ok {
		return result
	}
	return d
}

// XYZ2lonlat 行列号转经纬度
func XYZ2lonlat(x int, y int, z int) [2]float64 {
	n := math.Pow(2, float64(z))
	lonDeg := (float64(x)/n)*360.0 - 180.0
	latRad := math.Atan(math.Sinh(math.Pi * (1 - float64(2*y)/n)))
	latDeg := (180 * latRad) / math.Pi
	return [2]float64{lonDeg, latDeg}
}
