package storage

import (
	"crypto/rand"
	"fmt"
)

func GetId(prefix string) (string, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s_%x", prefix, b), nil
}
