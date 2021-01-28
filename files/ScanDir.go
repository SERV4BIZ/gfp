package files

import (
	"io/ioutil"
	"strings"

	"github.com/SERV4BIZ/gfp/handler"
)

// ScanDir is scan all file in path directory
func ScanDir(pathDir string) ([]string, error) {
	newPathDir := strings.TrimSpace(pathDir)
	mtx := MutexFile(newPathDir)
	mtx.RLock()
	defer mtx.RUnlock()

	files, err := ioutil.ReadDir(newPathDir)
	if !handler.Error(err) {
		length := len(files)
		filenames := make([]string, length)
		for i, name := range files {
			filenames[i] = name.Name()
		}
		return filenames, err
	}

	return make([]string, 0), err
}
