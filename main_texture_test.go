// TestInsides
package main

import (
	"math/rand"
	"testing"
	"tools_sdl/objects"
	"tools_sdl/structs"
	"tools_sdl/utils"

	"github.com/veandco/go-sdl2/sdl"
)

var TextureFiles map[string]string

const TEST_WIDTH = 800
const TEST_HEIGHT = 500

func TestTexture(t *testing.T) {

	textureFiles := make(map[string]string)
	textureFiles["GoIcon"] = "GoIcon.bmp"
	textureFiles["Dino"] = "dinosaur-walk-animated.gif"

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		TEST_WIDTH, TEST_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	tools.InitScaler()
	tools.InitPalette()

	textureManager, err := tools.LoadTextures(renderer, "", textureFiles)
	if err != nil {
		panic(err)
	}
	defer textureManager.TextureDestroy()

	rect1 := objects.NewRectangle("r1", -50, -50, 50, -50, 50, 50, -50, 50, 100, 100, tools.GetOpaqueColour("Black", 100), true)
	rect1.SetMovementData(structs.MovementData{Rotation: 80, X: 5, Y: 3})
	rect1.SetTextureData(textureManager.Get("GoIcon"))
	running := true
	for running {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.MouseButtonEvent:
				println("Quit")
				running = false
				break
			}
		}

		rect1.Update(0.05)
		if rect1.Point().Y > 450 {
			running = false
		}
		renderer.SetDrawColor(0, 100, 0, 255)
		renderer.Clear()
		rect1.Draw(renderer)

		for i := 0; i < 10000; i++ {
			x := rand.Intn(TEST_WIDTH)
			y := rand.Intn(TEST_HEIGHT)

			if rect1.PointInside(float64(x), float64(y)) {
				renderer.SetDrawColor(255, 255, 255, 255)
			} else {
				renderer.SetDrawColor(255, 0, 0, 255)
			}
			renderer.DrawPoint(int32(x), int32(y))
		}

		renderer.Present()
	}

}
