package hash

import (
	"crypto/sha256"
	"fmt"
)

// SHA256 is encode buffer byte to sha256 string
func SHA256(buffer []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(buffer))
}
