package world

import (
	"fmt"
	"strings"
	"github.com/balwes/hook/cmd/graphics"
	"github.com/balwes/hook/cmd/util"
)

type EntityKind int

const (
	UnknownEntity  EntityKind = iota
	GuyEntity
	WallEntity
)

func EntityKindToString(kind EntityKind) string {
	switch kind {
		case GuyEntity:
			return "Guy"
		case WallEntity:
			return "Wall"
		default:
			return fmt.Sprintf("Unknown (%d)", kind)
	}
}

func StringToEntityKind(kind string) EntityKind {
	switch strings.ToLower(kind) {
		case "guy":
			return GuyEntity
		case "wall":
			return WallEntity
		default:
			return UnknownEntity
	}
}

type Entity struct {
	id      int32
	kind    EntityKind
	Sprite  *graphics.Sprite
	IsDestroyed  bool
}

func (e Entity) String() string {
    return fmt.Sprintf("<%d:%s (%f,%f)>", e.id, EntityKindToString(e.kind), e.Sprite.Y, e.Sprite.Y)
}

func NewGuy(x, y float32) *Entity {
	e := Entity{}
	e.id     = -1
	e.kind   = GuyEntity
	e.Sprite = graphics.NewSprite(util.GetTexture("assets/images/guy.png"), x, y)
	e.IsDestroyed = false
	return &e
}

func NewWall(x, y float32) *Entity {
	e := Entity{}
	e.id     = -1
	e.kind   = WallEntity
	e.Sprite = graphics.NewSprite(util.GetTexture("assets/images/wall.png"), x, y)
	e.IsDestroyed = false
	return &e
}
