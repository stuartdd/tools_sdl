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
	World            *structs.World
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

func (p *Pic) Draw() {
	if p.Enabled {
		if p.TextureData != nil {
			p.World.Renderer.Copy(
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
		X: int32(p.XOrigin - p.World.X),
		Y: int32(p.YOrigin - p.World.Y),
	}
}

func (p *Pic) Rect() *sdl.Rect {
	return &sdl.Rect{
		X: int32((p.XOrigin - p.World.X) - p.W/2),
		Y: int32((p.YOrigin - p.World.Y) - p.H/2),
		W: int32(p.W),
		H: int32(p.H),
	}
}

func (p *Pic) PointInsideBounds(x float64, y float64) bool {
	return false
}

func NewPic(world *structs.World, pxOrigin, pyOrigin, W, H float64, textureData *structs.TextureData, enabled bool) *Pic {
	if textureData == nil {
		panic("Texture data is not defined")
	}
	return &Pic{
		XOrigin:      pxOrigin,
		YOrigin:      pyOrigin,
		W:            W,
		H:            H,
		MovementData: nil,
		World:        world,
		TextureData:  textureData,
		Enabled:      enabled,
	}
}
