package util

import (
	"crypto/sha512"
	"fmt"
)

// Hash returns a hash of the given string.
func Hash(data string) string {
	sum := sha512.Sum512([]byte(data))
	return fmt.Sprintf("%x", sum)
}
