// Triangle
package objects

import (
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type Triangle struct {
	Name                   string
	X1, Y1, X2, Y2, X3, Y3 int32
	Col                    sdl.Color
	Enabled                bool
}

func (p *Triangle) Draw(renderer *sdl.Renderer) {
	if p.Enabled {
		gfx.FilledTrigonColor(renderer, p.X1, p.Y1, p.X2, p.Y2, p.X3, p.Y3, p.Col)
	}
}

func NewTriangle(name string, px1, py1, px2, py2, px3, py3 int32, col sdl.Color, enabled bool) *Triangle {
	return &Triangle{
		Name:    name,
		X1:      px1,
		Y1:      py1,
		X2:      px2,
		Y2:      py2,
		X3:      px3,
		Y3:      py3,
		Col:     col,
		Enabled: enabled,
	}
}
