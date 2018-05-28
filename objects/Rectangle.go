// Rectangle
package objects

import (
	"tools_sdl/tools"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type Rectangle struct {
	Name                           string
	X1, Y1, X2, Y2, X3, Y3, X4, Y4 float64
	XOrigin, YOrigin               float64
	RotationAccu                   float64
	Rotation                       int
	RotationSpeed                  float64
	Col                            sdl.Color
	Enabled                        bool
}

func (p *Rectangle) Rotate(rot float64) {
	rotF := p.RotationAccu + rot
	rotP := p.Rotation
	rotFInt := int(rotF)
	if rotFInt != rotP {
		p.X1, p.Y1 = tools.Rotate(p.X1, p.Y1, rotP-rotFInt)
		p.X2, p.Y2 = tools.Rotate(p.X2, p.Y2, rotP-rotFInt)
		p.X3, p.Y3 = tools.Rotate(p.X3, p.Y3, rotP-rotFInt)
		p.X4, p.Y4 = tools.Rotate(p.X4, p.Y4, rotP-rotFInt)
		p.Rotation = rotFInt
	}
	p.RotationAccu = rotF
}

func (p *Rectangle) SetRotationSpeed(rs float64) {
	p.RotationSpeed = rs
}

func (p *Rectangle) Update(seconds float64) {
	p.Rotate(seconds * p.RotationSpeed)
}

func (p *Rectangle) Draw(renderer *sdl.Renderer) {
	if p.Enabled {
		xo := p.XOrigin
		yo := p.YOrigin
		gfx.FilledPolygonColor(
			renderer,
			[]int16{int16(xo + p.X1), int16(xo + p.X2), int16(xo + p.X3), int16(xo + p.X4)},
			[]int16{int16(yo + p.Y1), int16(yo + p.Y2), int16(yo + p.Y3), int16(yo + p.Y4)},
			p.Col)
	}
}

func (p *Rectangle) PointInside(x float64, y float64) bool {
	if p.PointInsideBounds(x, y) {
		return pointInsideTriangle(x, y, p.XOrigin, p.YOrigin, p.X2, p.Y2, p.X3, p.Y3, p.X4, p.Y4) ||
			pointInsideTriangle(x, y, p.XOrigin, p.YOrigin, p.X1, p.Y1, p.X2, p.Y2, p.X4, p.Y4)
	}
	return false
}

func pointInsideTriangle(x, y, x0, y0, x1, y1, x2, y2, x3, y3 float64) bool {
	dx := (x - x0) - x3
	dy := (y - y0) - y3

	dx32 := x3 - x2
	dy23 := y2 - y3

	D := dy23*(x1-x3) + dx32*(y1-y3)
	s := dy23*dx + dx32*dy
	t := (y3-y1)*dx + (x1-x3)*dy

	if D < 0 {
		return s <= 0 && t <= 0 && s+t >= D
	}
	return s >= 0 && t >= 0 && s+t <= D
}

func (p *Rectangle) PointInsideBounds(x float64, y float64) bool {
	xA := x - p.XOrigin
	minX := min4(p.X1, p.X2, p.X3, p.X4)
	if xA < minX {
		return false
	}
	maxX := max4(p.X1, p.X2, p.X3, p.X4)
	if xA > maxX {
		return false
	}
	yA := y - p.YOrigin
	minY := min4(p.Y1, p.Y2, p.Y3, p.Y4)
	if yA < minY {
		return false
	}
	maxY := max4(p.Y1, p.Y2, p.Y3, p.Y4)
	if yA > maxY {
		return false
	}
	return true
}

func NewRectangle(name string, px1, py1, px2, py2, px3, py3, px4, py4, pxOrigin, pyOrigin float64, col sdl.Color, enabled bool) *Rectangle {
	return &Rectangle{
		Name:          name,
		X1:            px1,
		Y1:            py1,
		X2:            px2,
		Y2:            py2,
		X3:            px3,
		Y3:            py3,
		X4:            px4,
		Y4:            py4,
		XOrigin:       pxOrigin,
		YOrigin:       pyOrigin,
		Rotation:      0,
		RotationAccu:  0.0,
		RotationSpeed: 0,
		Col:           col,
		Enabled:       enabled,
	}
}

func min4(a, b, c, d float64) float64 {
	if (a < b) && (a < c) && (a < d) {
		return a
	}
	if (b < c) && (b < d) {
		return b
	}
	if c < d {
		return c
	}
	return d
}

func max4(a, b, c, d float64) float64 {
	if (a > b) && (a > c) && (a > d) {
		return a
	}
	if (b > c) && (b > d) {
		return b
	}
	if c > d {
		return c
	}
	return d
}
