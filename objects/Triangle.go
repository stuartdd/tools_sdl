// Triangle
package objects

import (
	"tools_sdl/tools"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type Triangle struct {
	Name                   string
	X1, Y1, X2, Y2, X3, Y3 float64
	XOrigin, YOrigin       float64
	Rotation               float64
	Col                    sdl.Color
	Enabled                bool
}

func (p *Triangle) Rotate(rot float64) {
	p.Rotation = rot
}

func (p *Triangle) Position(ox, oy float64) {
	p.XOrigin = ox
	p.YOrigin = oy
}

func (p *Triangle) Update(seconds float64) {
	p.Rotation = p.Rotation - (seconds * 360)
}

func (p *Triangle) DrawX(renderer *sdl.Renderer) {
	xo := p.XOrigin
	yo := p.YOrigin
	if p.Enabled {
		gfx.FilledTrigonColor(renderer, int32(xo+p.X1), int32(yo+p.Y1), int32(xo+p.X2), int32(yo+p.Y2), int32(xo+p.X3), int32(yo+p.Y3), p.Col)
	}
}

func (p *Triangle) Draw(renderer *sdl.Renderer) {
	xo := p.XOrigin
	yo := p.YOrigin
	px1, py1 := tools.Rotate(p.X1, p.Y1, p.Rotation)
	px2, py2 := tools.Rotate(p.X2, p.Y2, p.Rotation)
	px3, py3 := tools.Rotate(p.X3, p.Y3, p.Rotation)
	if p.Enabled {
		gfx.FilledTrigonColor(renderer, int32(xo+px1), int32(yo+py1), int32(xo+px2), int32(yo+py2), int32(xo+px3), int32(yo+py3), p.Col)
	}
}

func NewTriangle(name string, px1, py1, px2, py2, px3, py3, pxOrigin, pyOrigin float64, col sdl.Color, enabled bool) *Triangle {
	return &Triangle{
		Name:     name,
		X1:       px1,
		Y1:       py1,
		X2:       px2,
		Y2:       py2,
		X3:       px3,
		Y3:       py3,
		XOrigin:  pxOrigin,
		YOrigin:  pyOrigin,
		Rotation: 0,
		Col:      col,
		Enabled:  enabled,
	}
}
