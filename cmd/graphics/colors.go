package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"github.com/balwes/hook/cmd/math"
)

var (
	ColorBlack      = sdl.Color{0, 0, 0, 255}
	ColorWhite      = sdl.Color{255, 255, 255, 255}
	ColorLightGray  = sdl.Color{191, 191, 191, 255}
	ColorGray       = sdl.Color{127, 127, 127, 255}
	ColorDarkGray   = sdl.Color{63, 63, 63, 255}
	ColorBlue       = sdl.Color{0, 0, 255, 255}
	ColorNavy       = sdl.Color{0, 0, 128, 255}
	ColorRoyal      = sdl.Color{65, 105, 225, 255}
	ColorSlate      = sdl.Color{112, 128, 144, 255}
	ColorSky        = sdl.Color{135, 206, 235, 255}
	ColorCyan       = sdl.Color{0, 255, 255, 255}
	ColorTeal       = sdl.Color{0, 128, 128, 255}
	ColorGreen      = sdl.Color{0, 255, 0, 255}
	ColorChartreuse = sdl.Color{127, 255, 0, 255}
	ColorLime       = sdl.Color{50, 205, 50, 255}
	ColorForest     = sdl.Color{34, 139, 34, 255}
	ColorOlive      = sdl.Color{107, 142, 35, 255}
	ColorYellow     = sdl.Color{255, 255, 0, 255}
	ColorGold       = sdl.Color{255, 215, 0, 255}
	ColorGoldenrod  = sdl.Color{218, 165, 32, 255}
	ColorOrange     = sdl.Color{255, 165, 0, 255}
	ColorBrown      = sdl.Color{139, 69, 19, 255}
	ColorTan        = sdl.Color{210, 180, 140, 255}
	ColorFirebrick  = sdl.Color{178, 34, 34, 255}
	ColorRed        = sdl.Color{255, 0, 0, 255}
	ColorScarlet    = sdl.Color{255, 52, 28, 255}
	ColorCoral      = sdl.Color{255, 127, 80, 255}
	ColorSalmon     = sdl.Color{250, 128, 114, 255}
	ColorPink       = sdl.Color{255, 105, 180, 255}
	ColorMagenta    = sdl.Color{255, 0, 255, 255}
	ColorPurple     = sdl.Color{160, 32, 240, 255}
	ColorViolet     = sdl.Color{238, 130, 238, 255}
	ColorMaroon     = sdl.Color{176, 48, 96, 255}
)

func RandomColor() sdl.Color {
	r := uint8(math.Round(rand.Float32() * 255.0))
	g := uint8(math.Round(rand.Float32() * 255.0))
	b := uint8(math.Round(rand.Float32() * 255.0))
	return sdl.Color{r,g,b,255}
}
