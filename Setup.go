package main

import (
	"github.com/stuartdd/tools_sdl/objects"
	"github.com/stuartdd/tools_sdl/tools"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()
	renderer.SetScale(2.0, 2.0)
	//	surface, err := window.GetSurface()
	//	if err != nil {
	//		panic(err)
	//	}
	//	surface.FillRect(nil, 0)

	//	rect := sdl.Rect{0, 0, 200, 200}
	//	surface.FillRect(&rect, 0xffff0000)

	palette := new(tools.Palette)
	palette.Init()
	tri := objects.NewTriangle("t1", 100, 100, 200, 200, 300, 100, palette.CMap["White"], true)
	tri.Draw(renderer)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
		renderer.SetDrawColor(0, 100, 0, 255)
		renderer.Clear()
		tri.Draw(renderer)
		renderer.Present()
	}
}
