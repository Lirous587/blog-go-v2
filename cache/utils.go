package cache

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateSecureToken Generate a secure random token
func GenerateSecureToken() (string, error) {
	bytes := make([]byte, 64) // 64 bytes = 512 bits
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
