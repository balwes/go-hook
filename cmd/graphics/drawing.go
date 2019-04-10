package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/balwes/go-hook/cmd/math"
	"github.com/balwes/go-hook/cmd/util"
)

func DrawPoint(cam *math.Camera, p *math.Point, color sdl.Color, size float32) {
	x, y := cam.WorldToScreen(p.X, p.Y)
	size *= cam.GetZoom()
	ok := gfx.FilledCircleColor(cam.Renderer, x, y, math.Round(size), color)
	util.PanicIfFalse(ok)
}

func DrawCircle(cam *math.Camera, c *math.Circle, color sdl.Color, filled bool) {
	x, y := cam.WorldToScreen(c.X, c.Y)
	r := math.Round(c.Radius * cam.GetZoom())
	var ok bool
	if filled {
		ok = gfx.FilledCircleColor(cam.Renderer, x, y, r, color)
	} else {
		ok = gfx.AACircleColor(cam.Renderer, x, y, r, color)
	}
	util.PanicIfFalse(ok)
}

func DrawLine(cam *math.Camera, l *math.Line, color sdl.Color, thickness float32) {
	x1, y1 := cam.WorldToScreen(l.X1, l.Y1)
	x2, y2 := cam.WorldToScreen(l.X2, l.Y2)
	thickness *= cam.GetZoom()
	var ok bool
	if thickness == 1 {
		ok = gfx.LineColor(cam.Renderer, x1, y1, x2, y2, color)
	} else {
		ok = gfx.ThickLineColor(cam.Renderer, x1, y1, x2, y2, math.Round(thickness), color)
	}
	util.PanicIfFalse(ok)
}

func DrawTriangle(cam *math.Camera, t *math.Triangle, color sdl.Color, filled bool) {
	x1, y1 := cam.WorldToScreen(t.X1, t.Y1)
	x2, y2 := cam.WorldToScreen(t.X2, t.Y2)
	x3, y3 := cam.WorldToScreen(t.X3, t.Y3)
	var ok bool
	if filled {
		ok = gfx.FilledTrigonColor(cam.Renderer, x1, y1, x2, y2, x3, y3, color)
	} else {
		ok = gfx.TrigonColor(cam.Renderer, x1, y1, x2, y2, x3, y3, color)
	}
	util.PanicIfFalse(ok)
}

func DrawRect(cam *math.Camera, r *math.Rect, color sdl.Color, filled bool) {
	x1, y1 := cam.WorldToScreen(r.X,           r.Y)
	x2, y2 := cam.WorldToScreen(r.X + r.Width, r.Y + r.Height)
	var ok bool
	if filled {
		ok = gfx.BoxColor(cam.Renderer, x1, y1, x2, y2, color)
	} else {
		ok = gfx.RectangleColor(cam.Renderer, x1, y1, x2, y2, color)
	}
	util.PanicIfFalse(ok)
}

func DrawPolygon(cam *math.Camera, p *math.Polygon, color sdl.Color, filled bool) {
	vertexCount := len(p.Vertices)
	vx := make([]int16, vertexCount)
	vy := make([]int16, vertexCount)
	for i := 0; i < vertexCount; i++ {
		x, y := cam.WorldToScreen(p.Vertices[i].X, p.Vertices[i].Y)
		vx[i] = int16(x)
		vy[i] = int16(y)
	}
	var ok bool
	if filled {
		ok = gfx.FilledPolygonColor(cam.Renderer, vx, vy, color)
	} else {
		ok = gfx.PolygonColor(cam.Renderer, vx, vy, color)
	}
	util.PanicIfFalse(ok)
}

func DrawRoundedRect(cam *math.Camera, r *math.Rect, radius float32, color sdl.Color, filled bool) {
	x1, y1 := cam.WorldToScreen(r.X,           r.Y)
	x2, y2 := cam.WorldToScreen(r.X + r.Width, r.Y + r.Height)
	radius *= cam.GetZoom()
	if radius < 2 {
		radius = 2 // Otherwise the filling won't be drawn
	}
	var ok bool
	if filled {
		ok = gfx.RoundedBoxColor(cam.Renderer, x1, y1, x2, y2, math.Round(radius), color)
	} else {
		ok = gfx.RoundedRectangleColor(cam.Renderer, x1, y1, x2, y2, math.Round(radius), color)
	}
	util.PanicIfFalse(ok)
}

func DrawRoundedRectOutline(cam *math.Camera, r *math.Rect, radius float32, color sdl.Color, thickness float32, lightenInnerRect bool) {
	// r is the outer rectangle
	inner := &math.Rect{
		r.X + thickness,
		r.Y + thickness,
		r.Width - thickness*2,
		r.Height - thickness*2,
	}
	innerRadius := radius * inner.Width / r.Width
	factor := float32(0.75)
	if lightenInnerRect {
		factor += 1
	}
	innerColor := sdl.Color{
		uint8(math.Round(math.Clamp(float32(color.R) * factor, 0, 255))),
		uint8(math.Round(math.Clamp(float32(color.G) * factor, 0, 255))),
		uint8(math.Round(math.Clamp(float32(color.B) * factor, 0, 255))),
		uint8(math.Round(math.Clamp(float32(color.A) * factor, 0, 255))),
	}
	DrawRoundedRect(cam,     r,      radius,      color, true)
	DrawRoundedRect(cam, inner, innerRadius, innerColor, true)
}

type TextureRegion struct {
	Texture *sdl.Texture
	X int32
	Y int32
	Width int32
	Height int32
}

type Sprite struct {
	texture *sdl.Texture
	X float32
	Y float32
	textureWidth  int32
	textureHeight int32
	ScaleX float32
	ScaleY float32
}

func NewSprite(tex *sdl.Texture, x, y float32) *Sprite {
	s := Sprite{}
	s.texture = tex
	s.X = x
	s.Y = y
	_,_,w,h,err := tex.Query()
	util.PanicIfNotNil(err)
	s.textureWidth = w
	s.textureHeight = h
	s.ScaleX = 1
	s.ScaleY = 1
	return &s
}

func (s *Sprite) Width() float32 {
	return float32(s.textureWidth) * s.ScaleX
}

func (s *Sprite) Height() float32 {
	return float32(s.textureHeight) * s.ScaleY
}

func (s Sprite) Draw(cam *math.Camera) {
	x, y := cam.WorldToScreen(s.X, s.Y)
	w := math.Round(float32( s.textureWidth) * cam.GetZoom() * s.ScaleX)
	h := math.Round(float32(s.textureHeight) * cam.GetZoom() * s.ScaleY)
	src := sdl.Rect{0, 0, s.textureWidth, s.textureHeight}
	dst := sdl.Rect{x, y, w, h}
	util.PanicIfNotNil(cam.Renderer.Copy(s.texture, &src, &dst))
}
