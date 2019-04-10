package math

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type Camera struct {
	X        float32 // @TODO rename these XOrigin or something.
	Y        float32 // @TODO rename these YOrigin or something.
	MinZoom  float32
	MaxZoom  float32
	zoom     float32
	Renderer *sdl.Renderer
}

func NewCamera(x, y float32, renderer *sdl.Renderer) *Camera {
	c := Camera{}
	c.X        = x
	c.Y        = y
	c.MinZoom  = 0.2
	c.MaxZoom  = 5
	c.zoom     = 1
	c.Renderer = renderer
	return &c
}

func (cam *Camera) GetZoom() float32 {
	return cam.zoom
}

func (cam *Camera) SetZoom(zoom float32) {
	cam.zoom = Clamp(zoom, cam.MinZoom, cam.MaxZoom)
}

func (cam *Camera) WorldToScreen(x, y float32) (int32, int32) {
	screenX := Round(x * cam.GetZoom() + cam.X)
	screenY := Round(y * cam.GetZoom() + cam.Y)
	return screenX, screenY
}

func (cam *Camera) ScreenToWorld(x, y int32) (float32, float32) {
	worldX := (float32(x) - cam.X) / cam.GetZoom()
	worldY := (float32(y) - cam.Y) / cam.GetZoom()
	return worldX, worldY
}

func (cam *Camera) ZoomTowards(zoomAmount, x, y float32) {
	oldZoom := cam.zoom
	cam.SetZoom(cam.zoom + zoomAmount)
	halfScreenWidth  := float32(1280/2)
	halfScreenHeight := float32( 720/2)
	dx := ( x -  halfScreenWidth) * (oldZoom - cam.zoom)
	dy := (-y + halfScreenHeight) * (oldZoom - cam.zoom)
	cam.X += dx
	cam.Y += dy
	log.Println("zoom =",cam.zoom,"x =",cam.X,"y =",cam.Y)
}
