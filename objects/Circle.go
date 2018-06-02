// Circle
package objects

import (
	"tools_sdl/structs"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type Circle struct {
	Name             string
	Radius           float64
	XOrigin, YOrigin float64
	MovementData     *structs.MovementData
	Col              sdl.Color
	Enabled          bool
	Filled           bool
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
	if p.MovementData != nil {
		p.XOrigin += p.MovementData.X
		p.YOrigin += p.MovementData.Y
	}
}

func (p *Circle) Draw(renderer *sdl.Renderer) {
	if p.Enabled {
		if p.Filled {
			gfx.FilledCircleColor(renderer, int32(p.XOrigin), int32(p.YOrigin), int32(p.Radius), p.Col)
		} else {
			gfx.CircleColor(renderer, int32(p.XOrigin), int32(p.YOrigin), int32(p.Radius), p.Col)
		}
	}
}

func (p *Circle) SetMovementData(md *structs.MovementData) {
	p.MovementData = md
}

func (p *Circle) GetMovementData() *structs.MovementData {
	return p.MovementData
}

func (p *Circle) SetColor(newCol sdl.Color) {
	p.Col = newCol
}

func NewCircle(name string, radius, pxOrigin, pyOrigin float64, col sdl.Color, enabled bool, filled bool) *Circle {
	return &Circle{
		Name:         name,
		XOrigin:      pxOrigin,
		YOrigin:      pyOrigin,
		Radius:       radius,
		MovementData: nil,
		Col:          col,
		Enabled:      enabled,
		Filled:       filled,
	}
}
