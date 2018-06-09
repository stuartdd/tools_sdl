package main

import (
	"fmt"
	"time"
	config "tools_jsonconfig"
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
		ImageLib:            "images",
		ApplicationDataPath: "../",
		KeyBindings:         make(map[string]string),
		TextureFiles:        make(map[string]string),
		DebugKeyboard:       true,
		FullScreen:          false,
		NonFsWidth:          800,
		NonFsHeight:         500,
		FsWidth:             1920,
		FsHeight:            1080,
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
			configData.FsWidth, configData.FsHeight, sdl.WINDOW_FULLSCREEN)
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

	windowW, windowH, err := renderer.GetOutputSize()
	if err != nil {
		panic(err)
	}
	mainWindowMiddleWidth := float64(windowW / 2)
	mainWindowMiddleHeight := float64(windowH / 2)

	world := &structs.World{Renderer: renderer, X: 0, Y: 0}

	textureManager, err := utils.LoadTextures(renderer, configData.ImageLib, configData.TextureFiles)
	if err != nil {
		panic(err)
	}
	defer textureManager.TextureDestroy()

	objects.InitScaler()
	utils.InitPalette()

	stars1 := objects.NewPicBasic(world, 0, 0, textureManager.Get("Stars_L_01"), true)
	lem := objects.NewRectangle(world, -60, -60, 60, -60, 60, 60, -60, 60, 400, 400, utils.GetOpaqueColour("Black", 100), true, true)
	lem.SetMovementData(&structs.MovementData{Rotation: 80, X: 0.01, Y: 0.03})
	lem.SetTextureData(textureManager.Get("LEM"))

	tri1 := objects.NewTriangle(world, -50, -50, 0, 51, 50, -50, 400, 300, utils.GetOpaqueColour("Coral Blue", 100), true, true)
	tri1.SetMovementData(&structs.MovementData{Rotation: 360 / 60, X: 0, Y: 0})
	cir1 := objects.NewCircle(world, 50, 400, 300, utils.GetOpaqueColour("Black", 100), true, true)

	collection := objects.NewCollection()
	collection.Add(lem)
	collection.Add(tri1)
	collection.Add(cir1)
	collection.Add(stars1)

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

		/*
			Track the LEM
		*/
		world.X = lem.XOrigin - mainWindowMiddleWidth
		world.Y = lem.YOrigin - mainWindowMiddleHeight

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		collection.UpdateDraw(timeDiff)

		renderer.Present()
	}
}
