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

func (p *Circle) PointInside(x float64, y float64) bool {
	if p.PointInsideBounds(x, y) {
		dx := x - p.XOrigin
		dy := y - p.YOrigin
		return (dx*dx+dy*dy <= p.Radius*p.Radius)
	}
	return false
}

func (p *Circle) PointInsideBounds(x float64, y float64) bool {
	if x < (p.XOrigin - p.Radius) {
		return false
	}
	if x > (p.XOrigin + p.Radius) {
		return false
	}

	if y < (p.YOrigin - p.Radius) {
		return false
	}
	if y > (p.YOrigin + p.Radius) {
		return false
	}
	return true

}

func (p *Circle) Update(seconds float64) {

}

func (p *Circle) Draw(renderer *sdl.Renderer) {
	if p.Enabled {
		gfx.FilledCircleColor(renderer, int32(p.XOrigin), int32(p.YOrigin), int32(p.Radius), p.Col)
	}
}

func (p *Circle) SetColor(newCol sdl.Color) {
	p.Col = newCol
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
