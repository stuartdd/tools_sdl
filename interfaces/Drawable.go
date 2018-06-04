// Drawable
package interfaces

type Drawable interface {
	Draw()
	Update(currentTime float64)
	PointInside(x float64, y float64) bool
	PointInsideBounds(x float64, y float64) bool
}
