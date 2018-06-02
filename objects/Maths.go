// Scaler
package objects

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

func PointInsideTriangle(x, y, x0, y0, x1, y1, x2, y2, x3, y3 float64) bool {
	dx := (x - x0) - x3
	dy := (y - y0) - y3

	dx32 := x3 - x2
	dy23 := y2 - y3

	D := dy23*(x1-x3) + dx32*(y1-y3)
	s := dy23*dx + dx32*dy
	t := (y3-y1)*dx + (x1-x3)*dy

	if D < 0 {
		return s <= 0 && t <= 0 && s+t >= D
	}
	return s >= 0 && t >= 0 && s+t <= D
}

func Min3(a, b, c float64) float64 {
	if (a < b) && (a < c) {
		return a
	}
	if b < c {
		return b
	}
	return c
}

func Max3(a, b, c float64) float64 {
	if (a > b) && (a > c) {
		return a
	}
	if b > c {
		return b
	}
	return c
}

func Min4(a, b, c, d float64) float64 {
	if (a < b) && (a < c) && (a < d) {
		return a
	}
	if (b < c) && (b < d) {
		return b
	}
	if c < d {
		return c
	}
	return d
}

func Max4(a, b, c, d float64) float64 {
	if (a > b) && (a > c) && (a > d) {
		return a
	}
	if (b > c) && (b > d) {
		return b
	}
	if c > d {
		return c
	}
	return d
}
