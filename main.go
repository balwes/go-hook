package main

//https://godoc.org/github.com/veandco/go-sdl2

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
	"log"
	"github.com/balwes/go-hook/cmd/world"
	"github.com/balwes/go-hook/cmd/math"
	"github.com/balwes/go-hook/cmd/util"
	"github.com/balwes/go-hook/cmd/graphics"
)

var window     *sdl.Window
var worldCam   *math.Camera
var hudCam     *math.Camera
var gameWorld  *world.World
var running = true
var line *math.Line
var rect *math.Rect
var lineClickerCounter = 0

func main() {
	log.Println("Hello")
	defer log.Println("Goodbye")
	err := sdl.Init(sdl.INIT_EVERYTHING)
	util.PanicIfNotNil(err)
	defer sdl.Quit() // Also quits subsystems

	// "nearest", "linear", "best"
	//sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "nearest") 

	window, err = sdl.CreateWindow(
		"Hook",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		1280,
		720,
		sdl.WINDOW_SHOWN ^ sdl.WINDOW_RESIZABLE,
	)
	util.PanicIfNotNil(err)
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	util.PanicIfNotNil(err)
	defer renderer.Destroy()

	util.InitCatalog(renderer)
	defer util.DestroyCatalog()
	
	ww, wh := window.GetSize()
	worldCam = math.NewCamera(float32(ww)/2, float32(wh)/2, renderer)
	hudCam = math.NewCamera(0, 0, renderer)
	
	math.GameWindow = window

	tmxPath := "assets/maps/map.tmx"
	gameWorld = world.NewWorld(tmxPath)
	util.WatchFile(tmxPath, func(path string) {
		gameWorld = world.NewWorld(tmxPath)
	})

	rect = &math.Rect{300,300,100,150}

	var before time.Time
	now := time.Now()
	// Should this be used for setting frame rate?
	// https://godoc.org/github.com/veandco/go-sdl2/gfx
	for running {
		before = now
		now = time.Now()
		dt := float32(now.Sub(before).Nanoseconds()) / 1000000000.0
		handleEvents(dt)
		update(dt)
		draw(dt)
		sdl.Delay(1)
	}
}

func handleEvents(dt float32) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			//case *sdl.KeyboardEvent:
			//	if t.Keysym.Sym == sdl.K_LEFT {
			//		fmt.Println("Pressed left")
			//	}
			case *sdl.MouseButtonEvent:
				if t.Button == sdl.BUTTON_LEFT && t.State == sdl.PRESSED {
					if lineClickerCounter == 0 {
						x, y := worldCam.ScreenToWorld(t.X, t.Y)
						line = &math.Line{x, y, x, y}
						lineClickerCounter++
					} else if lineClickerCounter == 1 {
						line = nil
						lineClickerCounter = 0
					}
				}
			case *sdl.MouseMotionEvent:
				if t.State == sdl.BUTTON_MIDDLE && (t.XRel != 0 || t.YRel != 0) {
					worldCam.X += float32(t.XRel)
					worldCam.Y += float32(t.YRel)
				}
				if lineClickerCounter == 1 {
					x, y := worldCam.ScreenToWorld(t.X, t.Y)
					line.X2 = x
					line.Y2 = y
				}
			case *sdl.MouseWheelEvent:
				if t.Y != 0 {
					zoomAmount := float32(t.Y) * 0.1
					worldCam.SetZoom(worldCam.GetZoom() + zoomAmount)
					//sx, sy, _ := sdl.GetMouseState()
					//wx, wy := worldCam.ScreenToWorld(sx, sy)
					//worldCam.ZoomTowards(zoomAmount, wx, wy)
				}
		}
	}
	if len(gameWorld.GetEntitiesByKind(world.GuyEntity)) > 0 {
		kb := sdl.GetKeyboardState()
		var dx float32 = 0
		var dy float32 = 0
		if kb[sdl.SCANCODE_LEFT] == 1 {
			dx -= 1
		}
		if kb[sdl.SCANCODE_RIGHT] == 1 {
			dx += 1
		}
		if kb[sdl.SCANCODE_UP] == 1 {
			dy -= 1
		}
		if kb[sdl.SCANCODE_DOWN] == 1 {
			dy += 1
		}
		gameWorld.GetEntitiesByKind(world.GuyEntity)[0].Sprite.X += dx*50*dt
		gameWorld.GetEntitiesByKind(world.GuyEntity)[0].Sprite.Y += dy*50*dt
	}
}

func update(dt float32) {
	gameWorld.Update(dt)
}

func draw(dt float32) {
	bg := graphics.ColorSky
	worldCam.Renderer.SetDrawColor(bg.R, bg.G, bg.B, 255)
	util.PanicIfNotNil(worldCam.Renderer.Clear())
	//
	gameWorld.Draw(worldCam)
	//t := &Triangle{0,0,10,10,10,0}
	//DrawTriangle(camera, t, ColorBlue, true)
	//c := &Circle{600,300,30}
	//DrawCircle(camera, c, ColorOlive, true)
	//DrawRect(camera, rect, ColorBlack, false)
	//if line != nil {
	//	DrawLine(camera, line, ColorSlate, 2)
	//}
	//DrawRoundedRectWithOutline(camera, &Rect{500,0,100,100}, 10, ColorGray, 10, false)
	//
	worldCam.Renderer.Present()
}
