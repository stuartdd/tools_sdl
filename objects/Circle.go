// Circle
package objects

import (
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type Circle struct {
	Name             string
	Radius           float64
	XOrigin, YOrigin float64
	Col              sdl.Color
	Enabled          bool
}

func (p *Circle) Update(seconds float64) {

}

func (p *Circle) Draw(renderer *sdl.Renderer) {
	if p.Enabled {
		gfx.FilledCircleColor(renderer, int32(p.XOrigin), int32(p.YOrigin), int32(p.Radius), p.Col)
	}
}
func NewCircle(name string, radius, pxOrigin, pyOrigin float64, col sdl.Color, enabled bool) *Circle {
	return &Circle{
		Name:    name,
		XOrigin: pxOrigin,
		YOrigin: pyOrigin,
		Radius:  radius,
		Col:     col,
		Enabled: enabled,
	}
}
