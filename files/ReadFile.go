package files

import (
	"io/ioutil"
	"strings"
)

// ReadFile is read file to buffer bytes from path file
func ReadFile(pathFile string) ([]byte, error) {
	newPathFile := strings.TrimSpace(pathFile)
	mtx := MutexFile(newPathFile)
	mtx.RLock()
	defer mtx.RUnlock()

	buffer, err := ioutil.ReadFile(newPathFile)
	if err != nil {
		return nil, err
	}
	return buffer, err
}
