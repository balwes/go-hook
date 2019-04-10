package editor

import (
	"log"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/balwes/go-hook/cmd/math"
	"github.com/balwes/go-hook/cmd/graphics"
	"github.com/balwes/go-hook/cmd/util"
	"github.com/balwes/go-hook/cmd/world"
)

var entityKindSelector  *graphics.ButtonGroup
var currentEntityKind   world.EntityKind

func InitEditor() {
	entityKindSelector = graphics.NewButtonGroup()

	{ // Guy entity selector button
		const ButtonSize = 50
		ButtonColor := graphics.ColorGoldenrod

		makeAndAddEntityKindSelectorButton := func(x, y float32, texPath string, kind world.EntityKind) {
			b := graphics.NewButton(x, y, ButtonSize, ButtonSize, ButtonColor)
			s := graphics.NewSprite(util.GetTexture(texPath), 0, 0)
			s.ScaleX = 0.25
			s.ScaleY = 0.25
			b.SetAndCenterSprite(s)
			b.OnClick = func() {
				currentEntityKind = kind
			}
			entityKindSelector.Add(b)
		}

		ww, _ := GameWindow.GetSize()
		const BeginY  = 200
		const Spacing = 5
		const MarginX = ButtonSize*2 + Spacing*2
		// Row 1
		makeAndAddEntityKindSelectorButton(
			float32(ww) - MarginX,
			BeginY,
			"assets/images/guy.png",
			world.GuyEntity)
		makeAndAddEntityKindSelectorButton(
			float32(ww) - MarginX + ButtonSize + Spacing,
			BeginY,
			"assets/images/crate.png",
			world.CrateEntity)
		// Row 2
		makeAndAddEntityKindSelectorButton(
			float32(ww) - MarginX,
			BeginY + ButtonSize + Spacing,
			"assets/images/dirt.png",
			world.DirtEntity)
		makeAndAddEntityKindSelectorButton(
			float32(ww) - MarginX + ButtonSize + Spacing,
			BeginY + ButtonSize + Spacing,
			"assets/images/pot.png",
			world.PotEntity)
		// Row 3
		makeAndAddEntityKindSelectorButton(
			float32(ww) - MarginX,
			BeginY + 2*ButtonSize + 2*Spacing,
			"assets/images/steel.png",
			world.SteelEntity)
	}
	currentEntityKind = world.UnknownEntity
}

func HandleEvent(event sdl.Event) {
	switch t := event.(type) {
		case *sdl.MouseButtonEvent:
			if t.Button == sdl.BUTTON_LEFT && t.State == sdl.PRESSED {
				mx, my := HudCam.ScreenToWorld(t.X, t.Y)
				clicked := entityKindSelector.TrySelect(mx, my)
				if !clicked {
					currentEntityKind = world.UnknownEntity
				}
			}
		//case *sdl.MouseMotionEvent:
		//	for _,b := range buttons {
		//		mouseWorldX, mouseWorldY := HudCam.ScreenToWorld(t.X, t.Y)
		//		p := &math.Point{mouseWorldX, mouseWorldY}
		//		b.IsHoveredOver = math.PointInsideRect(p, &b.Rect)
		//	}
	}
}

func Update(dt float32) {
	log.Printf(world.EntityKindToString(currentEntityKind))
}

func Draw(worldCam *math.Camera, dt float32) {
	entityKindSelector.Draw(HudCam)
	rect  := &math.Rect{0, 0, 100, 100}
	color := sdl.Color{20,180,70,255}
	graphics.DrawRoundedRectOutline(worldCam, rect, 3, color, 3, true)
}
