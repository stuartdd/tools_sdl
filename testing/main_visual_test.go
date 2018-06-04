// TestInsides
package testing

import (
	"math/rand"
	"testing"
	"tools_sdl/objects"
	"tools_sdl/structs"
	"tools_sdl/utils"

	"github.com/veandco/go-sdl2/sdl"
)

const WIDTH = 900
const HEIGHT = 600

func TestVisual(t *testing.T) {

	textureFiles := make(map[string]string)
	textureFiles["GoIcon"] = "GoIcon.bmp"

	t.Log("START:")
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		900, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	world := &structs.World{Renderer: renderer, X: 0, Y: 0}

	textureManager, err := utils.LoadTextures(renderer, "", textureFiles)
	if err != nil {
		panic(err)
	}
	defer textureManager.TextureDestroy()

	objects.InitScaler()
	utils.InitPalette()

	tri1 := objects.NewTriangle(world, -50, -50, 0, 51, 50, -50, 150, 100, utils.GetOpaqueColour("Black", 100), true, true)
	tri2 := objects.NewTriangle(world, -50, -50, 0, 51, 50, -50, 300, 100, utils.GetOpaqueColour("Black", 100), true, true)
	tri3 := objects.NewTriangle(world, -50, -50, 0, 51, 50, -50, 150, 250, utils.GetOpaqueColour("Black", 100), true, true)
	tri4 := objects.NewTriangle(world, -50, -50, 0, 51, 50, -50, 300, 250, utils.GetOpaqueColour("Black", 100), true, false)
	tri5 := objects.NewTriangle(world, 0, 0, 0, 50, 50, 50, 150, 400, utils.GetOpaqueColour("Black", 100), true, true)
	cir1 := objects.NewCircle(world, 50, 450, 100, utils.GetOpaqueColour("Black", 100), true, true)
	cir2 := objects.NewCircle(world, 50, 600, 100, utils.GetOpaqueColour("Black", 100), true, false)
	rect1 := objects.NewRectangle(world, -40, -40, 50, -50, 60, 60, -50, 50, 450, 250, utils.GetOpaqueColour("Black", 100), true, true)
	rect2 := objects.NewRectangle(world, -40, -40, 50, -50, 60, 60, -50, 50, 600, 250, utils.GetOpaqueColour("Black", 100), true, false)

	pic := objects.NewPic(world, 300, 400, 100, 100, textureManager.Get("GoIcon"), true)
	window.UpdateSurface()

	movement := structs.MovementData{Rotation: 60, X: 0, Y: 0}

	tri3.SetMovementData(&movement)
	tri4.SetMovementData(&movement)
	tri5.SetMovementData(&movement)
	rect1.SetMovementData(&movement)
	rect2.SetMovementData(&movement)
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
		world.X = world.X - 0.5

		renderer.SetDrawColor(0, 100, 0, 255)
		renderer.Clear()

		tri3.Update(0.05)
		tri4.Update(0.05)
		tri5.Update(0.05)
		rect1.Update(0.05)
		rect2.Update(0.05)
		tri1.Draw()
		tri2.Draw()
		tri3.Draw()
		tri4.Draw()
		tri5.Draw()
		cir1.Draw()
		cir2.Draw()
		rect1.Draw()
		rect2.Draw()
		pic.Draw()

		for i := 0; i < 10000; i++ {
			x := rand.Intn(WIDTH)
			y := rand.Intn(HEIGHT)

			if tri1.PointInsideBounds(float64(x), float64(y)) {
				renderer.SetDrawColor(255, 255, 255, 255)
			} else {
				if tri2.PointInside(float64(x), float64(y)) {
					renderer.SetDrawColor(255, 255, 255, 255)
				} else {
					if cir1.PointInsideBounds(float64(x), float64(y)) {
						renderer.SetDrawColor(255, 255, 255, 255)
					} else {
						if cir2.PointInside(float64(x), float64(y)) {
							renderer.SetDrawColor(255, 255, 255, 255)
						} else {
							if tri3.PointInsideBounds(float64(x), float64(y)) {
								renderer.SetDrawColor(255, 255, 255, 255)
							} else {
								if tri4.PointInside(float64(x), float64(y)) {
									renderer.SetDrawColor(255, 255, 255, 255)
								} else {
									if tri5.PointInside(float64(x), float64(y)) {
										renderer.SetDrawColor(255, 255, 255, 255)
									} else {
										if rect1.PointInsideBounds(float64(x), float64(y)) {
											renderer.SetDrawColor(255, 255, 255, 255)
										} else {
											if rect2.PointInside(float64(x), float64(y)) {
												renderer.SetDrawColor(255, 255, 255, 255)
											} else {
												renderer.SetDrawColor(255, 0, 0, 255)
											}
										}
									}
								}
							}
						}
					}
				}
			}
			renderer.DrawPoint(int32(x), int32(y))
		}

		renderer.Present()
	}

}
