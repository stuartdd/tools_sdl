// Scaler
package tools

import "math"

const DEG float64 = (2 * math.Pi) / 360

var SinTables [360]float64
var CosTables [360]float64

func RotateXY(px, py, ox, oy, decimalDegrees float64) (float64, float64) {
	norm := int(decimalDegrees) % 360
	if norm < 0 {
		norm = norm + 360
	}
	cosTheta := CosTables[norm]
	sinTheta := SinTables[norm]
	dx := cosTheta*(px-ox) - sinTheta*(py-oy) + ox
	dy := sinTheta*(px-ox) + cosTheta*(py-oy) + oy
	return dx, dy
}

func Rotate(px, py, decimalDegrees float64) (float64, float64) {
	norm := int(decimalDegrees) % 360
	if norm < 0 {
		norm = norm + 360
	}
	cosTheta := CosTables[norm]
	sinTheta := SinTables[norm]
	dx := cosTheta*px - sinTheta*py
	dy := sinTheta*px + cosTheta*py
	return dx, dy
}

func InitScaler() {
	var r float64 = 0
	for i := 0; i < 360; i++ {
		SinTables[i] = math.Sin(r)
		CosTables[i] = math.Cos(r)
		r += DEG
	}
}
