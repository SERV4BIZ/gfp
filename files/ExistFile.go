package files

import (
	"os"
	"strings"
)

// ExistFile is check exist file from pathfile
func ExistFile(pathFile string) bool {
	newPathFile := strings.TrimSpace(pathFile)
	mtx := MutexFile(newPathFile)
	mtx.RLock()
	defer mtx.RUnlock()

	_, err := os.Stat(newPathFile)
	if err != nil {
		return !os.IsNotExist(err)
	}
	return true
}
