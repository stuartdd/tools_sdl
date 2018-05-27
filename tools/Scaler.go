// Scaler
package tools

import "math"

const CIRCLE = 360
const DEG float64 = (2 * math.Pi) / CIRCLE

var SinTables [CIRCLE]float64
var CosTables [CIRCLE]float64

func RotateXY(px, py, ox, oy float64, degrees int) (float64, float64) {
	norm := degrees % CIRCLE
	if norm < 0 {
		norm = norm + CIRCLE
	}
	cosTheta := CosTables[norm]
	sinTheta := SinTables[norm]
	dx := cosTheta*(px-ox) - sinTheta*(py-oy) + ox
	dy := sinTheta*(px-ox) + cosTheta*(py-oy) + oy
	return dx, dy
}

func Rotate(px, py float64, degrees int) (float64, float64) {
	norm := int(degrees) % CIRCLE
	if norm < 0 {
		norm = norm + CIRCLE
	}
	cosTheta := CosTables[norm]
	sinTheta := SinTables[norm]
	dx := cosTheta*px - sinTheta*py
	dy := sinTheta*px + cosTheta*py
	return dx, dy
}

func InitScaler() {
	var r float64 = 0
	for i := 0; i < CIRCLE; i++ {
		SinTables[i] = math.Sin(r)
		CosTables[i] = math.Cos(r)
		r += DEG
	}
}
