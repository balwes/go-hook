package editor

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/balwes/go-hook/cmd/math"
)

var GameWindow  *sdl.Window
var HudCam      *math.Camera
var WorldCam    *math.Camera
