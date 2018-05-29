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

const WIDTH = 900
const HEIGHT = 600

func TestVisual(t *testing.T) {
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

	tools.InitScaler()
	tools.InitPalette()

	tri1 := objects.NewTriangle("t1", -50, -50, 0, 51, 50, -50, 150, 100, tools.GetOpaqueColour("Black", 100), true)
	tri2 := objects.NewTriangle("t2", -50, -50, 0, 51, 50, -50, 300, 100, tools.GetOpaqueColour("Black", 100), true)
	tri3 := objects.NewTriangle("t3", -50, -50, 0, 51, 50, -50, 150, 250, tools.GetOpaqueColour("Black", 100), true)
	tri4 := objects.NewTriangle("t4", -50, -50, 0, 51, 50, -50, 300, 250, tools.GetOpaqueColour("Black", 100), true)
	tri5 := objects.NewTriangle("t5", 0, 0, 0, 50, 50, 50, 150, 400, tools.GetOpaqueColour("Black", 100), true)
	cir1 := objects.NewCircle("c1", 50, 450, 100, tools.GetOpaqueColour("Black", 100), true)
	cir2 := objects.NewCircle("c2", 50, 600, 100, tools.GetOpaqueColour("Black", 100), true)
	rect1 := objects.NewRectangle("r1", -40, -40, 50, -50, 60, 60, -50, 50, 450, 250, tools.GetOpaqueColour("Black", 100), true)
	rect2 := objects.NewRectangle("r2", -40, -40, 50, -50, 60, 60, -50, 50, 600, 250, tools.GetOpaqueColour("Black", 100), true)
	window.UpdateSurface()

	tri3.SetMovementData(structs.MovementData{Rotation: 60, X: 0, Y: 0})
	tri4.SetMovementData(structs.MovementData{Rotation: 60, X: 0, Y: 0})
	tri5.SetMovementData(structs.MovementData{Rotation: 60, X: 0, Y: 0})
	rect1.SetMovementData(structs.MovementData{Rotation: 60, X: 0, Y: 0})
	rect2.SetMovementData(structs.MovementData{Rotation: 60, X: 0, Y: 0})
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
		tri5.Update(0.05)
		rect1.Update(0.05)
		rect2.Update(0.05)
		tri1.Draw(renderer)
		tri2.Draw(renderer)
		tri3.Draw(renderer)
		tri4.Draw(renderer)
		tri5.Draw(renderer)
		cir1.Draw(renderer)
		cir2.Draw(renderer)
		rect1.Draw(renderer)
		rect2.Draw(renderer)

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
