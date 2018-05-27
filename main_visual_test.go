// TestInsides
package main

import (
	"math/rand"
	"testing"
	"tools_sdl/objects"
	"tools_sdl/tools"

	"github.com/veandco/go-sdl2/sdl"
)

const WIDTH = 900
const HEIGHT = 600

func TestVisual(t *testing.T) {
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

	tools.InitScaler()
	tools.InitPalette()

	tri1 := objects.NewTriangle("t1", -50, -50, 0, 51, 50, -50, 150, 100, tools.GetOpaqueColour("Black", 100), true)
	tri2 := objects.NewTriangle("t1", -50, -50, 0, 51, 50, -50, 300, 100, tools.GetOpaqueColour("Black", 100), true)
	tri3 := objects.NewTriangle("t1", -50, -50, 0, 51, 50, -50, 150, 300, tools.GetOpaqueColour("Black", 100), true)
	tri4 := objects.NewTriangle("t1", -50, -50, 0, 51, 50, -50, 300, 300, tools.GetOpaqueColour("Black", 100), true)
	cir1 := objects.NewCircle("t1", 50, 450, 100, tools.GetOpaqueColour("Black", 100), true)
	cir2 := objects.NewCircle("t1", 50, 600, 100, tools.GetOpaqueColour("Black", 100), true)
	window.UpdateSurface()

	tri3.SetRotationSpeed(5)
	tri4.SetRotationSpeed(5)
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

		renderer.SetDrawColor(0, 100, 0, 255)
		renderer.Clear()

		tri3.Update(0.05)
		tri4.Update(0.05)
		tri1.Draw(renderer)
		tri2.Draw(renderer)
		tri3.Draw(renderer)
		tri4.Draw(renderer)
		cir1.Draw(renderer)
		cir2.Draw(renderer)

		for i := 0; i < 20000; i++ {
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
									renderer.SetDrawColor(255, 0, 0, 255)
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
