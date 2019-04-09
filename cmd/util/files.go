package util

import (
	"path/filepath"
	"log"
	"time"
	"os"
)

func GetFnameWithoutExtensionFromPath(path string) string {
	_, fname := filepath.Split(path)
	ext := filepath.Ext(fname)
	return fname[0:len(fname)-len(ext)]
}

func WatchFile(path string, f func(string)) {
	pause, err := time.ParseDuration("500ms")
	PanicIfNotNil(err)
	var lastTime time.Time
	firstIter := true
	go func() {
		for {
			finfo, err := os.Lstat(path)
			if (err != nil) {
				log.Printf("Could not open watched file \"%s\"\n", path)
				return
			}
			if firstIter {
				lastTime = finfo.ModTime()
				firstIter = false
			}
			if finfo.ModTime() != lastTime {
				f(path)
				lastTime = finfo.ModTime()
			}
			time.Sleep(pause)
		}
	}()
}

func ListFilesInDirectory(root string) []FileInfoWithPath {
	files := []FileInfoWithPath{}
	foo := func(path string, info os.FileInfo, err error) error {
		PanicIfNotNil(err)
		if !info.IsDir() {
			files = append(files, FileInfoWithPath{info, path})
		}
		return nil
	}
	err := filepath.Walk(root, foo)
	PanicIfNotNil(err)
	return files
}
