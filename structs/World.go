// World.go
package structs

import (
	"github.com/veandco/go-sdl2/sdl"
)

type World struct {
	Renderer *sdl.Renderer
	X        float64
	Y        float64
}
