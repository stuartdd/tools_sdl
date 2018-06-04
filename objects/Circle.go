// Circle
package objects

import (
	"tools_sdl/structs"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type Circle struct {
	Radius           float64
	XOrigin, YOrigin float64
	MovementData     *structs.MovementData
	World            *structs.World
	Col              sdl.Color
	Enabled          bool
	Filled           bool
}

func (p *Circle) PointInside(x float64, y float64) bool {
	if p.PointInsideBounds(x, y) {
		dx := x - (p.XOrigin - p.World.X)
		dy := y - (p.YOrigin - p.World.Y)
		return (dx*dx+dy*dy <= p.Radius*p.Radius)
	}
	return false
}

func (p *Circle) PointInsideBounds(x float64, y float64) bool {
	if x < ((p.XOrigin - p.World.X) - p.Radius) {
		return false
	}
	if x > ((p.XOrigin - p.World.X) + p.Radius) {
		return false
	}
	if y < ((p.YOrigin - p.World.Y) - p.Radius) {
		return false
	}
	if y > ((p.YOrigin - p.World.Y) + p.Radius) {
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

func (p *Circle) Draw() {
	if p.Enabled {
		xo := p.XOrigin - p.World.X
		yo := p.YOrigin - p.World.Y
		if p.Filled {
			gfx.FilledCircleColor(p.World.Renderer, int32(xo), int32(yo), int32(p.Radius), p.Col)
		} else {
			gfx.CircleColor(p.World.Renderer, int32(xo), int32(yo), int32(p.Radius), p.Col)
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

func NewCircle(world *structs.World, radius, pxOrigin, pyOrigin float64, col sdl.Color, enabled bool, filled bool) *Circle {
	return &Circle{
		XOrigin:      pxOrigin,
		YOrigin:      pyOrigin,
		Radius:       radius,
		MovementData: nil,
		World:        world,
		Col:          col,
		Enabled:      enabled,
		Filled:       filled,
	}
}
