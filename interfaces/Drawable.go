// Drawable
package interfaces

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Drawable interface {
	Draw(renderer *sdl.Renderer)
	Update(currentTime float64)
	PointInside(x float64, y float64) bool
	PointInsideBounds(x float64, y float64) bool
}
