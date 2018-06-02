// TestInsides
package testing

import (
	"fmt"
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

	objects.InitScaler()
	utils.InitPalette()

	textureManager, err := utils.LoadTextures(renderer, "", textureFiles)
	if err != nil {
		panic(err)
	}
	defer textureManager.TextureDestroy()

	if textureManager.Get("GoIcon") == nil {
		panic(fmt.Errorf("GoIcon Not loaded"))
	}

	rect1 := objects.NewRectangle("r1", -50, -50, 50, -50, 50, 50, -50, 50, 100, 100, utils.GetOpaqueColour("Black", 100), true, true)
	rect1.SetMovementData(&structs.MovementData{Rotation: 15, X: 1, Y: 1})
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

		rect1.Update(0.1)
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
