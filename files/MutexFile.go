package files

import (
	"strings"
	"sync"
)

// MutexMapMutexFile is Mutex lock MapMutexFile variable
var MutexMapMutexFile sync.RWMutex

// MapMutexFile is data map for mutex lock path file
var MapMutexFile = make(map[string]*sync.RWMutex)

// MutexFile is get mutex lock path file from keyname
func MutexFile(key string) *sync.RWMutex {
	keyName := strings.TrimSpace(key)
	MutexMapMutexFile.RLock()
	mtx, ok := MapMutexFile[keyName]
	if ok {
		MutexMapMutexFile.RUnlock()
		return mtx
	}
	MutexMapMutexFile.RUnlock()

	MutexMapMutexFile.Lock()
	mtx = new(sync.RWMutex)
	MapMutexFile[keyName] = mtx
	MutexMapMutexFile.Unlock()
	return mtx
}
