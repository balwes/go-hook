package editor

import (
	"log"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/balwes/hook/cmd/math"
	"github.com/balwes/hook/cmd/graphics"
)

func HandleEvent(event sdl.Event) {
	switch t := event.(type) {
		case *sdl.MouseButtonEvent:
			if t.Button == sdl.BUTTON_LEFT && t.State == sdl.PRESSED {
				log.Println("editor recognized mouse left button press")
			}
	}
}

func Update(dt float32) {
}

func Draw(worldCam *math.Camera, hudCam *math.Camera, dt float32) {
	//ww, wh := GameWindow.GetSize()
	rect  := &math.Rect{0, 0, 100, 100}//float32(ww)/2-250-10, -float32(wh)/2+100, 250, 500}
	color := sdl.Color{20,180,70,255}
	graphics.DrawRoundedRectWithOutline(worldCam, rect, 3, color, 3, true)
	graphics.DrawRoundedRectWithOutline(  hudCam, rect, 3, sdl.Color{255,255,255,255}, 3, false)
}
