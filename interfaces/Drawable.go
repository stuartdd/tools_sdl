// Drawable
package interfaces

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Drawable interface {
	Draw(renderer *sdl.Renderer)
	Update(currentTime int64)
}
