package files

import (
	"os"
	"strings"
)

// IsDir is check path file it is a directory
func IsDir(pathFile string) bool {
	newPathFile := strings.TrimSpace(pathFile)
	mtx := MutexFile(newPathFile)
	mtx.RLock()
	defer mtx.RUnlock()

	pf, err := os.Stat(newPathFile)
	if err != nil {
		return !os.IsNotExist(err)
	}
	return pf.Mode().IsDir()
}
