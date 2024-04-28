package vault

import (
	"log"
	"os"
	"path/filepath"
)

func createSecretAndPacks(root string) {
	err := os.MkdirAll(root, os.ModePerm)
	if err != nil {
		log.Fatalf("Couldn't create directory: %s\n", err.Error())
	}

	secretFile := filepath.Join(root, secretsFileName)
	_, err = os.Create(secretFile)
	if err != nil {
		log.Fatalf("Error creating a secret file: %s\n", err.Error())
	}

	packsFile := filepath.Join(root, "packs")
	err = os.MkdirAll(packsFile, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating packs: %s\n", err.Error())
	}
}

func writeSecret(root string, secret string) {
	fD, err := os.OpenFile(
		filepath.Join(root, secretsFileName),
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		log.Fatalf("Error opening secrets: %s\n", err.Error())
	}
	defer fD.Close()

	fD.WriteString(secret)
}

func GetProjectName() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Cannot stat: %s\n", err.Error())
	}
	absPath := filepath.Dir(pwd)
	dir := filepath.Base(absPath)
	return dir
}
