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
	RotationAccu           float64
	Rotation               int
	RotationSpeed          float64
	Col                    sdl.Color
	Enabled                bool
}

func (p *Triangle) Rotate(rot float64) {
	rotF := p.RotationAccu + rot
	rotP := p.Rotation
	rotFInt := int(rotF)
	if rotFInt != rotP {
		p.X1, p.Y1 = tools.Rotate(p.X1, p.Y1, rotP-rotFInt)
		p.X2, p.Y2 = tools.Rotate(p.X2, p.Y2, rotP-rotFInt)
		p.X3, p.Y3 = tools.Rotate(p.X3, p.Y3, rotP-rotFInt)
		p.Rotation = rotFInt
	}
	p.RotationAccu = rotF
}

func (p *Triangle) Update(seconds float64) {
	p.Rotate(seconds * p.RotationSpeed)
}

func (p *Triangle) Position(ox, oy float64) {
	p.XOrigin = ox
	p.YOrigin = oy
}

func (p *Triangle) Draw(renderer *sdl.Renderer) {
	if p.Enabled {
		xo := p.XOrigin
		yo := p.YOrigin
		gfx.FilledTrigonColor(renderer, int32(xo+p.X1), int32(yo+p.Y1), int32(xo+p.X2), int32(yo+p.Y2), int32(xo+p.X3), int32(yo+p.Y3), p.Col)
	}
}

func (p *Triangle) SetColor(newCol sdl.Color) {
	p.Col = newCol
}

func (p *Triangle) SetRotationSpeed(rs float64) {
	p.RotationSpeed = rs
}

func (p *Triangle) PointInside(x float64, y float64) bool {
	if p.PointInsideBounds(x, y) {

		dx := (x - p.XOrigin) - p.X3
		dy := (y - p.YOrigin) - p.Y3

		dx32 := p.X3 - p.X2
		dy23 := p.Y2 - p.Y3

		D := dy23*(p.X1-p.X3) + dx32*(p.Y1-p.Y3)
		s := dy23*dx + dx32*dy
		t := (p.Y3-p.Y1)*dx + (p.X1-p.X3)*dy

		if D < 0 {
			return s <= 0 && t <= 0 && s+t >= D
		}
		return s >= 0 && t >= 0 && s+t <= D
	}
	return false
}

func (p *Triangle) PointInsideBounds(x float64, y float64) bool {
	xA := x - p.XOrigin
	minX := min(p.X1, p.X2, p.X3)
	if xA < minX {
		return false
	}
	maxX := max(p.X1, p.X2, p.X3)
	if xA > maxX {
		return false
	}
	yA := y - p.YOrigin
	minY := min(p.Y1, p.Y2, p.Y3)
	if yA < minY {
		return false
	}
	maxY := max(p.Y1, p.Y2, p.Y3)
	if yA > maxY {
		return false
	}
	return true
}

func NewTriangle(name string, px1, py1, px2, py2, px3, py3, pxOrigin, pyOrigin float64, col sdl.Color, enabled bool) *Triangle {
	return &Triangle{
		Name:          name,
		X1:            px1,
		Y1:            py1,
		X2:            px2,
		Y2:            py2,
		X3:            px3,
		Y3:            py3,
		XOrigin:       pxOrigin,
		YOrigin:       pyOrigin,
		Rotation:      0,
		RotationAccu:  0.0,
		RotationSpeed: 0,
		Col:           col,
		Enabled:       enabled,
	}
}

func min(a float64, b float64, c float64) float64 {
	if (a < b) && (a < c) {
		return a
	}
	if b < c {
		return b
	}
	return c
}

func max(a float64, b float64, c float64) float64 {
	if (a > b) && (a > c) {
		return a
	}
	if b > c {
		return b
	}
	return c
}
