package world

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"path/filepath"
	"github.com/balwes/go-hook/cmd/math"
	"github.com/balwes/go-hook/cmd/graphics"
	"github.com/balwes/go-hook/cmd/util"
	"github.com/lafriks/go-tiled"
)

const TileSize  = 32
const TileSizeF = float32(TileSize)

type World struct {
	Path string
	tiledMap *tiled.Map
	tmxTileCache map[uint32]graphics.TextureRegion
	entities []*Entity
	entityCounter int32
}

func NewWorld(path string) *World {
	tmap, err := tiled.LoadFromFile(path)
	util.PanicIfNotNil(err)
	w := World{}
	w.Path = path
	w.tiledMap = tmap
	w.tmxTileCache = map[uint32]graphics.TextureRegion{}
	w.entities = []*Entity{}
	w.entityCounter = 0
	w.addObjectsFromTmx()
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
	world.drawTiledMap(cam)
	for _,e := range world.entities {
		e.Sprite.Draw(cam)
	}
}

// Adapted from https://github.com/lafriks/go-tiled/blob/master/render/renderer.go
func (w *World) drawTiledMap(cam *math.Camera) {
	m := w.tiledMap
	for _, layer := range m.Layers {
		if m.RenderOrder != "right-down" {
			log.Fatalf("Tmx \"%s\" has an unsupported render order\n", w.Path)
		}

		i := 0
		for r := 0; r < m.Height; r++ {
			for c := 0; c < m.Width; c++ {
				if layer.Tiles[i].IsNil() {
					i++
					continue
				}
				if layer.Opacity != 1 {
					log.Fatalf("Bad Tmx layer opacity\n")
				}

				img, err := w.getTileImage(layer.Tiles[i])
				util.PanicIfNotNil(err)
	
				src := sdl.Rect{img.X, img.Y, img.Width, img.Height}
				z := cam.GetZoom()
				dst := sdl.Rect{
					math.Round(float32(c * m.TileWidth)  * z + cam.X),
					math.Round(float32(r * m.TileHeight) * z + cam.Y),
					math.Round(float32(img.Width)  * z + 0.5), // +0.5 to prevent gaps
					math.Round(float32(img.Height) * z + 0.5), // +0.5 to prevent gaps
				}
				cam.Renderer.CopyEx(img.Texture,
									&src,
									&dst,
								    0,
							        nil,
									sdl.FLIP_NONE)
				i++
			}
		}
	}
}

// Adapted from https://github.com/lafriks/go-tiled/blob/master/render/renderer.go
func (w *World) getTileImage(tile *tiled.LayerTile) (graphics.TextureRegion, error) {
	timg, ok := w.tmxTileCache[tile.Tileset.FirstGID+tile.ID]
	// Precache all tiles in tileset
	if !ok {
		mapFileParentDir,_ := filepath.Split(w.Path)
		path := filepath.Join(mapFileParentDir, tile.Tileset.Image.Source)
		img := util.GetTexture(path)

		tilesetColumns := tile.Tileset.Columns
		if tilesetColumns == 0 {
			tilesetColumns = tile.Tileset.Image.Width / (tile.Tileset.TileWidth + tile.Tileset.Spacing)
		}

		tilesetTileCount := tile.Tileset.TileCount
		if tilesetTileCount == 0 {
			tilesetTileCount = (tile.Tileset.Image.Height / (tile.Tileset.TileHeight + tile.Tileset.Spacing)) * tilesetColumns
		}

		for i := tile.Tileset.FirstGID; i < tile.Tileset.FirstGID + uint32(tilesetTileCount); i++ {
			x := int(i-tile.Tileset.FirstGID) % tilesetColumns
			y := int(i-tile.Tileset.FirstGID) / tilesetColumns

			w.tmxTileCache[i] = graphics.TextureRegion{
				img,
				int32(x * tile.Tileset.TileWidth),
				int32(y * tile.Tileset.TileHeight),
				int32(tile.Tileset.TileWidth),
				int32(tile.Tileset.TileHeight),
			}
			if tile.ID == i-tile.Tileset.FirstGID {
				timg = w.tmxTileCache[i]
			}
		}
	}

	return timg, nil
}

func (world *World) addObjectsFromTmx() {
	for _, group := range world.tiledMap.ObjectGroups {
		// @TODO Visible is always false
		//if !group.Visible {
		//	continue
		//}
		for _, object := range group.Objects {
			// @TODO Visible is always false
			//if !object.Visible {
			//	continue
			//}
			x := float32(object.X)
			y := float32(object.Y)
			w := float32(object.Width)
			h := float32(object.Height)
			kind := StringToEntityKind(object.Type)

			if kind == UnknownEntity {
				log.Printf("Unknown entity at (%f,%f) won't be added\n", x, y)
			} else if !math.FloatIsWhole(x / TileSizeF) || !math.FloatIsWhole(y / TileSizeF) {
					log.Printf("%s with bad position (%f,%f) won't be added (must be a multiple of %d)\n", object.Type, x, y, TileSize)
			} else if w != TileSizeF || h != TileSizeF {
				log.Printf("%s with bad size at (%f,%f) won't be added (must be (%d,%d))\n", object.Type, x, y, TileSize, TileSize)
			} else {
				e := NewEntity(x, y, kind)
				e.Sprite.ScaleX = (w / TileSizeF) * (w / e.Sprite. Width())
				e.Sprite.ScaleY = (h / TileSizeF) * (h / e.Sprite.Height())
				e.Sprite.Y -= h // Why?
				world.AddEntity(e)
			}
		}
	}
}
