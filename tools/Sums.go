// Sums
package tools

func Min(a float64, b float64, c float64) float64 {
	if (a < b) && (a < c) {
		return a
	}
	if b < c {
		return b
	}
	return c
}

func Max(a float64, b float64, c float64) float64 {
	if (a > b) && (a > c) {
		return a
	}
	if b > c {
		return b
	}
	return c
}
