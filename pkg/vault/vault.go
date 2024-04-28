package vault

import (
	"log"
	"os"
)

func NewVault(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatalf("Couldn't create directory: %s\n", err.Error())
	}
}

func Pack(filePath string) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Couldn't read given env file: %s\n", err.Error())
	}
}
