package helpers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

const (
	salt = "KriFTWFd2rdUMXm"
)

// GenerateTokenWithSalt ...
func GenerateTokenWithSalt(text string) string {
	data := []byte(text + salt + time.Now().String())
	hash := fmt.Sprintf("%x", md5.Sum(data))
	return hash
}

// PasswordHashed ...
func PasswordHashed(plainPassword string) string {
	bytes := sha256.Sum256([]byte(plainPassword))
	hashedPassword := hex.EncodeToString(bytes[:])
	return hashedPassword
}
