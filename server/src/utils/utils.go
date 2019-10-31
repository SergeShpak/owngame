package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateToken(len int) (string, error) {
	t := make([]byte, len)
	if _, err := rand.Read(t); err != nil {
		return "", err
	}
	t64 := base64.RawURLEncoding.EncodeToString(t)
	return t64, nil
}
