package main

import (
	"time"

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

	tools.InitScaler()
	tools.InitPalette()

	tri1 := objects.NewTriangle("t1", -50, -50, 0, 51, 50, -50, 400, 300, tools.GetColour("Coral Blue"), true)
	cir1 := objects.NewCircle("t1", 50, 400, 300, tools.GetOpaqueColour("Black", 100), true)
	window.UpdateSurface()
	var timeLast int64 = time.Now().UnixNano()
	var timeDiff float64 = 0
	var timeTemp int64 = 0
	running := true
	for running {
		timeTemp = time.Now().UnixNano()
		timeDiff = float64(timeTemp-timeLast) / 1000000000.0
		timeLast = timeTemp

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
		tri1.Update(timeDiff)

		tri1.Draw(renderer)
		cir1.Draw(renderer)
		renderer.Present()
	}
}
