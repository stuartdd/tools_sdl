// Drawable
package interfaces

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Drawable interface {
	Draw(surface *sdl.Surface, x int32, y int32)
}
