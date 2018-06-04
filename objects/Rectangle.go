// Rectangle
package objects

import (
	"tools_sdl/structs"

	"github.com/veandco/go-sdl2/gfx"

	"github.com/veandco/go-sdl2/sdl"
)

type Rectangle struct {
	X1, Y1, X2, Y2, X3, Y3, X4, Y4 float64
	XOrigin, YOrigin               float64
	RotationAccu                   float64
	Rotation                       int
	MovementData                   *structs.MovementData
	TextureData                    *structs.TextureData
	World                          *structs.World
	Col                            sdl.Color
	Enabled                        bool
	Filled                         bool
	W                              int32
	H                              int32
}

func (p *Rectangle) rotate(rot float64) {
	if rot > 0.0 {
		rotF := p.RotationAccu + rot
		rotP := p.Rotation
		rotFInt := int(rotF)
		if rotFInt != rotP {
			p.X1, p.Y1 = Rotate(p.X1, p.Y1, rotP-rotFInt)
			p.X2, p.Y2 = Rotate(p.X2, p.Y2, rotP-rotFInt)
			p.X3, p.Y3 = Rotate(p.X3, p.Y3, rotP-rotFInt)
			p.X4, p.Y4 = Rotate(p.X4, p.Y4, rotP-rotFInt)
			p.Rotation = rotFInt
		}
		p.RotationAccu = rotF
	}
}

func (p *Rectangle) SetMovementData(md *structs.MovementData) {
	p.MovementData = md
}

func (p *Rectangle) GetMovementData() *structs.MovementData {
	return p.MovementData
}

func (p *Rectangle) SetTextureData(td *structs.TextureData) {
	p.TextureData = td
}

func (p *Rectangle) GetTextureData() *structs.TextureData {
	return p.TextureData
}

func (p *Rectangle) Update(seconds float64) {
	if p.MovementData != nil {
		p.rotate(seconds * p.MovementData.Rotation)
		p.XOrigin += p.MovementData.X
		p.YOrigin += p.MovementData.Y
	}
}

func (p *Rectangle) Draw() {
	if p.Enabled {
		if p.TextureData != nil {
			p.World.Renderer.CopyEx(p.TextureData.Texture, p.TextureData.Rect, p.Rect(), -p.RotationAccu, &sdl.Point{X: p.W / 2, Y: p.H / 2}, sdl.FLIP_NONE)
		} else {
			xo := p.XOrigin - p.World.X
			yo := p.YOrigin - p.World.Y
			if p.Filled {
				gfx.FilledPolygonColor(
					p.World.Renderer,
					[]int16{int16(xo + p.X1), int16(xo + p.X2), int16(xo + p.X3), int16(xo + p.X4)},
					[]int16{int16(yo + p.Y1), int16(yo + p.Y2), int16(yo + p.Y3), int16(yo + p.Y4)},
					p.Col)
			} else {
				gfx.PolygonColor(
					p.World.Renderer,
					[]int16{int16(xo + p.X1), int16(xo + p.X2), int16(xo + p.X3), int16(xo + p.X4)},
					[]int16{int16(yo + p.Y1), int16(yo + p.Y2), int16(yo + p.Y3), int16(yo + p.Y4)},
					p.Col)
			}
		}
	}
}

func (p *Rectangle) PointInside(x float64, y float64) bool {
	if p.PointInsideBounds(x, y) {
		xo := p.XOrigin - p.World.X
		yo := p.YOrigin - p.World.Y
		return PointInsideTriangle(x, y, xo, yo, p.X2, p.Y2, p.X3, p.Y3, p.X4, p.Y4) ||
			PointInsideTriangle(x, y, xo, yo, p.X1, p.Y1, p.X2, p.Y2, p.X4, p.Y4)
	}
	return false
}

func (p *Rectangle) PointInsideBounds(x float64, y float64) bool {
	xA := x - (p.XOrigin - p.World.X)
	minX := Min4(p.X1, p.X2, p.X3, p.X4)
	if xA < minX {
		return false
	}
	maxX := Max4(p.X1, p.X2, p.X3, p.X4)
	if xA > maxX {
		return false
	}
	yA := y - (p.YOrigin - p.World.Y)
	minY := Min4(p.Y1, p.Y2, p.Y3, p.Y4)
	if yA < minY {
		return false
	}
	maxY := Max4(p.Y1, p.Y2, p.Y3, p.Y4)
	if yA > maxY {
		return false
	}
	return true
}

func (p *Rectangle) Point() *sdl.Point {
	return &sdl.Point{
		X: int32(p.XOrigin - p.World.X),
		Y: int32(p.YOrigin - p.World.Y),
	}
}

func (p *Rectangle) Rect() *sdl.Rect {
	return &sdl.Rect{
		X: int32(p.XOrigin-p.World.X) - (p.W / 2),
		Y: int32(p.YOrigin-p.World.Y) - (p.H / 2),
		W: p.W,
		H: p.H,
	}
}

func NewRectangle(world *structs.World, px1, py1, px2, py2, px3, py3, px4, py4, pxOrigin, pyOrigin float64, col sdl.Color, enabled bool, filled bool) *Rectangle {
	X1 := int32(Min4(px1, px2, px3, px4))
	Y1 := int32(Min4(py1, py2, py3, py4))
	X2 := int32(Max4(px1, px2, px3, px4))
	Y2 := int32(Max4(py1, py2, py3, py4))
	W := X2 - X1
	H := Y2 - Y1

	return &Rectangle{
		World:        world,
		X1:           px1,
		Y1:           py1,
		X2:           px2,
		Y2:           py2,
		X3:           px3,
		Y3:           py3,
		X4:           px4,
		Y4:           py4,
		XOrigin:      pxOrigin,
		YOrigin:      pyOrigin,
		Rotation:     0,
		RotationAccu: 0.0,
		MovementData: nil,
		TextureData:  nil,
		Col:          col,
		Enabled:      enabled,
		Filled:       filled,
		W:            W,
		H:            H,
	}
}
