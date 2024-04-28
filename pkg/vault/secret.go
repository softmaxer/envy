package vault

import (
	"crypto/rand"
	"log"
)

func createSecret() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatalf("Error reading into key bytes: %s\n", err.Error())
	}
	return string(key)
}
