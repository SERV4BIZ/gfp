package files

import (
	"os"
)

// GetTempDir is get path temp directory from os
func GetTempDir() string {
	return os.TempDir()
}
