// Rectangle
package objects

import (
	"tools_sdl/structs"

	"github.com/veandco/go-sdl2/sdl"
)

type Pic struct {
	W, H             float64
	XOrigin, YOrigin float64
	MovementData     *structs.MovementData
	TextureData      *structs.TextureData
	Enabled          bool
}

func (p *Pic) SetMovementData(md *structs.MovementData) {
	p.MovementData = md
}

func (p *Pic) GetMovementData() *structs.MovementData {
	return p.MovementData
}

func (p *Pic) SetTextureData(td *structs.TextureData) {
	p.TextureData = td
}

func (p *Pic) GetTextureData() *structs.TextureData {
	return p.TextureData
}

func (p *Pic) Update(seconds float64) {
	if p.MovementData != nil {
		p.XOrigin += p.MovementData.X
		p.YOrigin += p.MovementData.Y
	}
}

func (p *Pic) Draw(renderer *sdl.Renderer) {
	if p.Enabled {
		if p.TextureData != nil {
			renderer.Copy(
				p.TextureData.Texture,
				p.TextureData.Rect,
				p.Rect())
		}
	}
}

func (p *Pic) PointInside(x float64, y float64) bool {
	return false
}

func (p *Pic) Point() *sdl.Point {
	return &sdl.Point{
		X: int32(p.XOrigin),
		Y: int32(p.YOrigin),
	}
}

func (p *Pic) Rect() *sdl.Rect {
	return &sdl.Rect{
		X: int32(p.XOrigin - p.W/2),
		Y: int32(p.YOrigin - p.H/2),
		W: int32(p.W),
		H: int32(p.H),
	}
}

func (p *Pic) PointInsideBounds(x float64, y float64) bool {
	return false
}

func NewPic(pxOrigin, pyOrigin, W, H float64, textureData *structs.TextureData, enabled bool) *Pic {

	return &Pic{
		XOrigin:     pxOrigin,
		YOrigin:     pyOrigin,
		W:           W,
		H:           H,
		TextureData: textureData,
		Enabled:     enabled,
	}
}
