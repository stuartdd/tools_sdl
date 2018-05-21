// Drawable
package interfaces

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Drawable interface {
	Draw(renderer *sdl.Renderer)
	Update(currentTime float64)
	InsideBounds(x float64, y float64) bool
}
