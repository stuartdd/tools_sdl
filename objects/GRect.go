package objects

import (
	"github.com/veandco/go-sdl2/sdl"
)

type GRect struct {
	Name    string
	Shape   *sdl.Rect
	Col     *sdl.Color
	Enabled bool
}

func NewGRect(name string, x int32, y int32, w int32, h int32, c1 sdl.Color) *GRect {
	return &GRect{
		Name: name,
		Shape: &sdl.Rect{
			X: x,
			Y: y,
			W: w,
			H: h,
		},
		Col: &c1,
	}
}

func (p *GRect) GetName() string {
	return p.Name
}

func (p *GRect) GetShape() *sdl.Rect {
	return p.Shape
}

func (p *GRect) Draw(renderer *sdl.Renderer) {
	if p.Enabled {
		renderer.SetDrawColor(p.Col.R, p.Col.G, p.Col.B, p.Col.A)
		renderer.FillRect(p.Shape)
	}
}

func (p *GRect) Move(x int32, y int32) {
	p.Shape.X = x
	p.Shape.Y = y
}
