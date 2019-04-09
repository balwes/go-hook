package util

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"
	"path/filepath"
	"fmt"
	"os"
	"strings"
)

const PathToAssets = "assets"
var theRenderer *sdl.Renderer
var textures = map[string]*sdl.Texture{} // Maps all the possible resource paths to the potentially loaded resource

type FileInfoWithPath struct {
    os.FileInfo
    Path string
}

func InitCatalog(renderer *sdl.Renderer) {
	theRenderer = renderer
	filenames := ListFilesInDirectory(PathToAssets)
	resourceCount := 0
	for _,f := range filenames {
		ext := strings.ToLower(filepath.Ext(f.Name()))
		if ext == ".png" {
			textures[f.Path] = nil // Not yet loaded
			resourceCount++
		} else {
			fmt.Printf("Catalog: Unknown resource \"%s\" will not be loadable\n", f.Path)
		}
	}
	fmt.Printf("Catalog: %d resource(s) found\n", resourceCount)
}

func DestroyCatalog() {
	destroyCount := 0
	for _,t := range textures {
		if t != nil {
			destroyCount++
			t.Destroy()
		}
	}
	fmt.Printf("Catalog: %d resource(s) destroyed\n", destroyCount)
}

func GetTexture(filepathWithExt string) *sdl.Texture {
	filepathWithExt = filepath.FromSlash(filepathWithExt)
	texture, found := textures[filepathWithExt]
	if !found {
		return nil
	}
	if texture == nil {
		// Load and cache the texture
		surf, err := img.Load(filepathWithExt)
		PanicIfNotNil(err)
		defer surf.Free()
		loadedTexture, err := theRenderer.CreateTextureFromSurface(surf)
		PanicIfNotNil(err)
		textures[filepathWithExt] = loadedTexture
		return loadedTexture
	} else {
		return texture
	}
}
