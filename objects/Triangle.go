// Triangle
package objects

import (
	"tools_sdl/structs"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type Triangle struct {
	Name                   string
	X1, Y1, X2, Y2, X3, Y3 float64
	XOrigin, YOrigin       float64
	RotationAccu           float64
	Rotation               int
	MovementData           *structs.MovementData
	Col                    sdl.Color
	Enabled                bool
	Filled                 bool
	W                      int32
	H                      int32
}

func (p *Triangle) rotate(rot float64) {
	rotF := p.RotationAccu + rot
	rotP := p.Rotation
	rotFInt := int(rotF)
	if rotFInt != rotP {
		p.X1, p.Y1 = Rotate(p.X1, p.Y1, rotP-rotFInt)
		p.X2, p.Y2 = Rotate(p.X2, p.Y2, rotP-rotFInt)
		p.X3, p.Y3 = Rotate(p.X3, p.Y3, rotP-rotFInt)
		p.Rotation = rotFInt
	}
	p.RotationAccu = rotF
}

func (p *Triangle) Update(seconds float64) {
	if p.MovementData != nil {
		p.rotate(seconds * p.MovementData.Rotation)
		p.XOrigin += p.MovementData.X
		p.YOrigin += p.MovementData.Y
	}
}

func (p *Triangle) Draw(renderer *sdl.Renderer) {
	if p.Enabled {
		if p.Filled {
			xo := p.XOrigin
			yo := p.YOrigin
			gfx.FilledTrigonColor(renderer, int32(xo+p.X1), int32(yo+p.Y1), int32(xo+p.X2), int32(yo+p.Y2), int32(xo+p.X3), int32(yo+p.Y3), p.Col)
		} else {
			xo := p.XOrigin
			yo := p.YOrigin
			gfx.TrigonColor(renderer, int32(xo+p.X1), int32(yo+p.Y1), int32(xo+p.X2), int32(yo+p.Y2), int32(xo+p.X3), int32(yo+p.Y3), p.Col)
		}
	}
}

func (p *Triangle) SetMovementData(md *structs.MovementData) {
	p.MovementData = md
}

func (p *Triangle) GetMovementData() *structs.MovementData {
	return p.MovementData
}

func (p *Triangle) SetColor(newCol sdl.Color) {
	p.Col = newCol
}

func (p *Triangle) Point() *sdl.Point {
	return &sdl.Point{
		X: int32(p.XOrigin),
		Y: int32(p.YOrigin),
	}
}

func (p *Triangle) Rect() *sdl.Rect {
	return &sdl.Rect{
		X: int32(p.XOrigin) - (p.W / 2),
		Y: int32(p.YOrigin) - (p.H / 2),
		W: p.W,
		H: p.H,
	}
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
	minX := Min3(p.X1, p.X2, p.X3)
	if xA < minX {
		return false
	}
	maxX := Max3(p.X1, p.X2, p.X3)
	if xA > maxX {
		return false
	}
	yA := y - p.YOrigin
	minY := Min3(p.Y1, p.Y2, p.Y3)
	if yA < minY {
		return false
	}
	maxY := Max3(p.Y1, p.Y2, p.Y3)
	if yA > maxY {
		return false
	}
	return true
}

func NewTriangle(name string, px1, py1, px2, py2, px3, py3, pxOrigin, pyOrigin float64, col sdl.Color, enabled bool, filled bool) *Triangle {
	X1 := int32(Min3(px1, px2, px3))
	Y1 := int32(Min3(py1, py2, py3))
	X2 := int32(Max3(px1, px2, px3))
	Y2 := int32(Max3(py1, py2, py3))
	W := X2 - X1
	H := Y2 - Y1

	return &Triangle{
		Name:         name,
		X1:           px1,
		Y1:           py1,
		X2:           px2,
		Y2:           py2,
		X3:           px3,
		Y3:           py3,
		XOrigin:      pxOrigin,
		YOrigin:      pyOrigin,
		Rotation:     0,
		RotationAccu: 0.0,
		MovementData: nil,
		Col:          col,
		Enabled:      enabled,
		Filled:       filled,
		W:            W,
		H:            H,
	}
}
