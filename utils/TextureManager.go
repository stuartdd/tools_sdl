// Textures
package utils

import (
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
		textureData, err := loadTexture(renderer, applicationDataPath+fileName, name)
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

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, err
	}
	return &structs.TextureData{Name: name,
		Texture: texture,
		Rect:    &surface.ClipRect,
	}, nil
}
