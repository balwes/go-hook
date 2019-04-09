package world

import (
	"log"
	"github.com/balwes/hook/cmd/math"
)

type World struct {
	Path string
	entities []*Entity
	entityCounter int32
}

func NewWorld(path string) *World {
	w := World{}
	w.Path = path
	w.entities = []*Entity{}
	w.entityCounter = 0
	return &w
}

func (w *World) AddEntity(e *Entity) {
	w.entities = append(w.entities, e)
	e.id = w.entityCounter
	w.entityCounter++
}

func (world *World) GetEntitiesByKind(kind EntityKind) []*Entity {
	entities := []*Entity{}
	for _,e := range world.entities {
		if e.kind == kind {
			entities = append(entities, e)
		}
	}
	return entities
}

func (world *World) Update(dt float32) {
	for _,e := range world.entities {
		switch e.kind {
			case GuyEntity:
				//
			case WallEntity:
				//
			default:
				log.Fatalf("Tried to update unknown entity %v\n", e)
		}
	}

	// Filter out dead entities
	// @TODO make this more efficient using this
	// https://stackoverflow.com/a/37335777
	entities := []*Entity{}
	for _,e := range world.entities {
		if !e.IsDestroyed {
			entities = append(entities, e)
		}
	}
	world.entities = entities
}

func (world *World) Draw(cam *math.Camera) {
	for _,e := range world.entities {
		switch e.kind {
			case GuyEntity:
				e.Sprite.Draw(cam)
			case WallEntity:
				e.Sprite.Draw(cam)
			default:
				log.Fatalf("Tried to draw unknown entity %v\n", e)
		}
	}
}
