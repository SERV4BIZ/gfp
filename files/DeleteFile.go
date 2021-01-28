package files

import (
	"os"
	"strings"
)

// DeleteFile is delete file from pathfile
func DeleteFile(pathFile string) error {
	newPathFile := strings.TrimSpace(pathFile)
	mtx := MutexFile(newPathFile)
	mtx.Lock()
	defer mtx.Unlock()

	err := os.RemoveAll(newPathFile)
	return err
}
