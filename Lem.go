package main

import (
	"fmt"
	"time"
	"tools-config"
	"tools_sdl/objects"
	"tools_sdl/structs"
	"tools_sdl/utils"

	"github.com/veandco/go-sdl2/sdl"
)

const NANO_PER_SECOND float64 = 1000000000

var configData structs.ConfigData
var mainWindow *sdl.Window
var timeLast int64 = time.Now().UnixNano()
var timeDiff float64 = 0
var timeTemp int64 = 0
var running bool = true

func main() {

	configData = structs.ConfigData{
		Name:                "Undefined",
		ApplicationDataPath: "../",
		KeyBindings:         make(map[string]string),
		TextureFiles:        make(map[string]string),
		DebugKeyboard:       true,
		FullScreen:          false,
		NonFsWidth:          800,
		NonFsHeight:         500,
	}

	err := config.LoadJson("config.json", &configData)
	if err != nil {
		panic(err)
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	if configData.FullScreen {
		window, err := sdl.CreateWindow(configData.Name, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			-1, -1, sdl.WINDOW_FULLSCREEN)
		if err != nil {
			panic(err)
		}
		mainWindow = window
	} else {
		window, err := sdl.CreateWindow(configData.Name, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			configData.NonFsWidth, configData.NonFsHeight, sdl.WINDOW_SHOWN)
		if err != nil {
			panic(err)
		}
		mainWindow = window
	}
	defer mainWindow.Destroy()

	renderer, err := sdl.CreateRenderer(mainWindow, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	textureManager, err := utils.LoadTextures(renderer, "", configData.TextureFiles)
	if err != nil {
		panic(err)
	}
	defer textureManager.TextureDestroy()

	objects.InitScaler()
	utils.InitPalette()

	rect1 := objects.NewRectangle("r1", -60, -60, 60, -60, 60, 60, -60, 60, 100, 100, utils.GetOpaqueColour("Black", 100), true, true)
	rect1.SetMovementData(&structs.MovementData{Rotation: 80, X: 0.2, Y: 0.1})
	rect1.SetTextureData(textureManager.Get("LEM"))

	tri1 := objects.NewTriangle("t1", -50, -50, 0, 51, 50, -50, 400, 300, utils.GetOpaqueColour("Coral Blue", 100), true, true)
	tri1.SetMovementData(&structs.MovementData{Rotation: 360 / 60, X: 0, Y: 0})
	cir1 := objects.NewCircle("t1", 50, 400, 300, utils.GetOpaqueColour("Black", 100), true, true)

	for running {
		timeTemp := time.Now().UnixNano()
		timeDiff = float64(timeTemp-timeLast) / NANO_PER_SECOND
		timeLast = timeTemp
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.MouseMotionEvent:
				if tri1.PointInside(float64(t.X), float64(t.Y)) {
					tri1.SetColor(utils.GetColour("Blue"))
				} else {
					tri1.SetColor(utils.GetColour("Coral Blue"))
				}
				break
			case *sdl.MouseButtonEvent:
				fmt.Printf("MouseButton %d %d\n", t.X, t.Y)
			case *sdl.KeyboardEvent:
				if t.State == sdl.PRESSED {
					keyName := sdl.GetKeyName(t.Keysym.Sym)
					mappedKey := configData.KeyBindings[keyName]
					if mappedKey == "" {
						mappedKey = keyName
					}
					if configData.DebugKeyboard {
						fmt.Printf("KeyboardEvent: KeyName is [%s] maps to [%s]\n", keyName, mappedKey)
					}
					if mappedKey == "Quit" {
						running = false
					}
				}
			}
		}

		renderer.SetDrawColor(0, 100, 0, 255)
		renderer.Clear()
		rect1.Update(timeDiff)
		tri1.Update(timeDiff)
		tri1.Draw(renderer)
		cir1.Draw(renderer)
		rect1.Draw(renderer)
		renderer.Present()
	}
}
