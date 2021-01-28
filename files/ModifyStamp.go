package files

import (
	"os"
	"strings"

	"github.com/SERV4BIZ/gfp/handler"
)

// ModifyStamp is get last modify unix stamp from path file
func ModifyStamp(pathFile string) (int64, error) {
	newPathFile := strings.TrimSpace(pathFile)
	mtx := MutexFile(newPathFile)
	mtx.RLock()
	defer mtx.RUnlock()

	pf, err := os.Stat(newPathFile)
	if handler.Error(err) {
		return -1, err
	}
	return pf.ModTime().Unix(), err
}
