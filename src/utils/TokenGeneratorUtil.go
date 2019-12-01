package utils

import (
	"crypto/rand"
	"fmt"
)

// Generate token
func TokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
