// TextureData.go
package structs

import (
	"github.com/veandco/go-sdl2/sdl"
)

type TextureData struct {
	Name    string
	Texture *sdl.Texture
	Rect    *sdl.Rect
}
