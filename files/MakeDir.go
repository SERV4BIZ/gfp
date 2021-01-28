package files

import (
	"os"
	"strings"
)

// MakeDir is create a new dirctory from path
func MakeDir(pathDir string) error {
	newPathDir := strings.TrimSpace(pathDir)
	mtx := MutexFile(newPathDir)
	mtx.Lock()
	defer mtx.Unlock()

	err := os.MkdirAll(newPathDir, 0777)
	return err
}
