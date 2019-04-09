package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"github.com/balwes/hook/cmd/util"
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
