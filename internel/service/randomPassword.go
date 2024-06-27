package service

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func GenerateRandomPassword(length int) (string, error) {
	byteLength := (length*6 + 7) / 8
	buffer := make([]byte, byteLength)

	_, err := io.ReadFull(rand.Reader, buffer)
	if err != nil {
		return "", err
	}

	password := base64.StdEncoding.EncodeToString(buffer)

	return password[:length], nil
}
