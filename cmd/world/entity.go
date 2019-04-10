package world

import (
	"fmt"
	"log"
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

func genericEntity() *Entity {
	e := Entity{}
	e.id     = -1
	e.kind   = UnknownEntity
	e.Sprite = nil
	e.IsDestroyed = false
	return &e
}

func NewEntity(x, y float32, kind EntityKind) *Entity {
	e := genericEntity()
	e.kind = kind
	texPath := ""
	switch kind {
		case GuyEntity:
			texPath = "assets/images/guy.png"
		case DirtEntity:
			texPath = "assets/images/dirt.png"
		case PotEntity:
			texPath = "assets/images/pot.png"
		case SteelEntity:
			texPath = "assets/images/steel.png"
		case CrateEntity:
			texPath = "assets/images/crate.png"
		default:
			log.Fatalf("Tried to create unknown entity %v\n", kind)
	}
	s := graphics.NewSprite(util.GetTexture(texPath), x, y)
	e.Sprite = s
	return e
}
