package world

import (
	"fmt"
	"github.com/balwes/go-hook/cmd/graphics"
	"github.com/balwes/go-hook/cmd/util"
)

type EntityKind int

const (
	UnknownEntity  EntityKind = iota
	GuyEntity
	DirtEntity
	PotEntity
	SteelEntity
	CrateEntity
)

func EntityKindToString(kind EntityKind) string {
	switch kind {
		case GuyEntity:
			return "Guy"
		case DirtEntity:
			return "Dirt"
		case PotEntity:
			return "Pot"
		case SteelEntity:
			return "Steel"
		case CrateEntity:
			return "Crate"
		default:
			return fmt.Sprintf("Unknown (%d)", kind)
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
