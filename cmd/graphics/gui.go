package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"github.com/balwes/go-hook/cmd/util"
	"github.com/balwes/go-hook/cmd/math"
)

type Label struct {
	text string
	font *ttf.Font
	x int32
	y int32
	texture *sdl.Texture
}

func NewLabel(text string, font *ttf.Font, x int32, y int32) *Label {
	l := new(Label)
	l.text = text
	l.font = font
	l.x = x
	l.y = y
    return l
}

func (l Label) Draw(renderer *sdl.Renderer) {
	// @TODO this doesn't use a camera
	blended, err := l.font.RenderUTF8Blended(l.text, sdl.Color{255,0,0,255})
	util.PanicIfNotNil(err)
	texture, err := renderer.CreateTextureFromSurface(blended)
	util.PanicIfNotNil(err)
	rect := sdl.Rect{}
	rect.X = l.x
	rect.Y = l.y
	rect.W = 40//l.Width()  // @TODO
	rect.H = 40//l.Height() // @TODO
	err = renderer.Copy(texture, nil, &rect)
	util.PanicIfNotNil(err)
}

func (l Label) Text() string {
	return l.text
}

func (l Label) Width() int32 {
	_, _, w, _, err := l.texture.Query()
	util.PanicIfNotNil(err)
	return w
}

func (l Label) Height() int32 {
	_, _, _, h, err := l.texture.Query()
	util.PanicIfNotNil(err)
	return h
}

type Button struct {
	Rect        math.Rect
	Color       sdl.Color
	Label       *Label
	Sprite      *Sprite
	IsSelected  bool
	OnClick     func()
}

func NewButton(x, y, w, h float32, color sdl.Color) *Button {
	b := &Button{}
	b.Rect       = math.Rect{x, y, w, h}
	b.Color      = color
	b.Label      = nil
	b.Sprite     = nil
	b.IsSelected = false
	b.OnClick    = func() {}
    return b
}

func (b *Button) SetAndCenterSprite(spr *Sprite) {
	rcx, rcy := b.Rect.Center()
	spr.X = rcx-spr.Width()/2
	spr.Y = rcy-spr.Height()/2
	b.Sprite = spr
}

func (b *Button) Draw(cam *math.Camera) {
	c := b.Color
	if b.IsSelected {
		c = MultiplyColor(c, 1.1)
	}
	DrawRect(cam, &b.Rect, c, true)
	if b.Label != nil {
		// @TODO Draw the label
	}
	if b.Sprite != nil {
		b.Sprite.Draw(cam)
	}
}

type ButtonGroup struct {
	Buttons []*Button
}

func NewButtonGroup() *ButtonGroup {
	bg := ButtonGroup{}
	bg.Buttons = []*Button{}
	return &bg
}

func (bg *ButtonGroup) Add(b *Button) {
	bg.Buttons = append(bg.Buttons, b)
}

func (bg *ButtonGroup) TrySelect(x, y float32) bool {
	clicked := false
	p := &math.Point{x,y}
	for _,b := range bg.Buttons {
		if math.PointInsideRect(p, &b.Rect) {
			b.OnClick()
			b.IsSelected = true
			clicked = true
		} else {
			b.IsSelected = false
		}
	}
	return clicked
}

func (bg *ButtonGroup) Draw(cam *math.Camera) {
	for _,b := range bg.Buttons {
		b.Draw(cam)
	}
}
