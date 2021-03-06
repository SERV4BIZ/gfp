package files

import (
	"io/ioutil"
	"strings"
)

// WriteFile is write buffer byte to path file
func WriteFile(pathFile string, buffer []byte) (int, error) {
	newPathFile := strings.TrimSpace(pathFile)
	mtx := MutexFile(newPathFile)
	mtx.Lock()
	defer mtx.Unlock()

	size := len(buffer)
	err := ioutil.WriteFile(newPathFile, buffer, 0777)
	if err != nil {
		return -1, err
	}
	return size, err
}
