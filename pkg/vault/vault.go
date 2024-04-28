package vault

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	secretsFileName string = "envy.secret"
)

func NewVault(root string) {
	createSecretAndPacks(root)
	vaultSecret := createSecret()
	writeSecret(root, vaultSecret)
}

func Pack(project string, envFile io.Reader) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Unable to access home: %s\n", err.Error())
	}
	packExt := fmt.Sprintf("%s.pack", project)
	packName := filepath.Join(homeDir, ".envy", "packs", packExt)
	packFD, err := os.Create(packName)
	if err != nil {
		log.Fatalf("Error creating file: %s\n", err.Error())
	}
	defer packFD.Close()
	keyName := filepath.Join(homeDir, ".envy", secretsFileName)
	keyFD, err := os.OpenFile(keyName, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("Error getting secret key: %s\n", err.Error())
	}
	defer keyFD.Close()
	hash, err := encrypt(keyFD, envFile)
	if err != nil {
		log.Fatalf("Error encoding the file: %s\n", err.Error())
	}
	packFD.Write(hash)
}

func Unpack(project string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Unable to access home: %s\n", err.Error())
	}
	packExt := fmt.Sprintf("%s.pack", project)
	packName := filepath.Join(homeDir, ".envy", "packs", packExt)
	packFD, err := os.Open(packName)
	if err != nil {
		log.Fatalf("Error opening pack: %s\n", err.Error())
	}
	defer packFD.Close()
	keyName := filepath.Join(homeDir, ".envy", secretsFileName)
	keyFD, err := os.OpenFile(keyName, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("Error getting secret key: %s\n", err.Error())
	}
	defer keyFD.Close()
	decodedEnv, err := decrypt(keyFD, packFD)
	fmt.Printf("%s", decodedEnv)
}
