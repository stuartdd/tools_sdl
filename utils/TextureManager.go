// Textures
package utils

import (
	"path/filepath"
	"tools_sdl/structs"

	"github.com/veandco/go-sdl2/sdl"
)

type TextureManager struct {
	Textures map[string]*structs.TextureData
}

func (p *TextureManager) TextureDestroy() {
	for _, t := range p.Textures {
		t.Texture.Destroy()
	}
}

func (p *TextureManager) Get(name string) *structs.TextureData {
	return p.Textures[name]
}

func LoadTextures(renderer *sdl.Renderer, applicationDataPath string, fileNames map[string]string) (*TextureManager, error) {
	list := make(map[string]*structs.TextureData)
	for name, fileName := range fileNames {
		var fn string
		if applicationDataPath == "" {
			fn = fileName
		} else {
			fn = filepath.Join(applicationDataPath, fileName)
		}

		textureData, err := loadTexture(renderer, fn, name)
		if err != nil {
			return nil, err
		}
		list[name] = textureData
	}
	return &TextureManager{
		Textures: list,
	}, nil
}

func loadTexture(renderer *sdl.Renderer, fileName string, name string) (*structs.TextureData, error) {
	surface, err := sdl.LoadBMP(fileName)
	if err != nil {
		return nil, err
	}
	defer surface.Free()
	cRect := &surface.ClipRect

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, err
	}
	/*
		Return the Texture data and a clone of it's rect.
		Had a bug where the rect was corrupted after a random time. I assume it
		was garbage collected when the defered serface.Free() ran as it was part
		of surface. Making a clone resolved the issue.
	*/
	return &structs.TextureData{Name: name,
		Texture: texture,
		Rect:    &sdl.Rect{X: cRect.X, Y: cRect.Y, W: cRect.W, H: cRect.H}, // the clone
	}, nil
}
