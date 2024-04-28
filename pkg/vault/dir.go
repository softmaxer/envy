package vault

import (
	"log"
	"os"
	"path/filepath"

	"github.com/softmaxer/envy/pkg/styles"
)

func createSecretAndPacks(root string) {
	err := os.MkdirAll(root, os.ModePerm)
	if err != nil {
		log.Fatal(styles.ErrorText().Render("Couldn't create directory: ", err.Error()))
	}

	secretFile := filepath.Join(root, secretsFileName)
	_, err = os.Create(secretFile)
	if err != nil {
		log.Fatal(styles.ErrorText().Render("Error creating a secret file: ", err.Error()))
	}

	packsFile := filepath.Join(root, "packs")
	err = os.MkdirAll(packsFile, os.ModePerm)
	if err != nil {
		log.Fatal(styles.ErrorText().Render("Error creating packs: ", err.Error()))
	}
}

func writeSecret(root string, secret string) {
	fD, err := os.OpenFile(
		filepath.Join(root, secretsFileName),
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		log.Fatal(styles.ErrorText().Render("Error opening secrets: ", err.Error()))
	}
	defer fD.Close()

	fD.WriteString(secret)
}

func GetProjectName() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(styles.ErrorText().Render("Cannot stat: ", err.Error()))
	}
	dir := filepath.Base(pwd)
	return dir
}

func WriteDecodedPack(decodedBytes []byte) {
	env, err := os.Create(".env")
	if err != nil {
		log.Fatal(styles.ErrorText().Render("Error creating a new .env file: ", err.Error()))
	}
	defer env.Close()
	env.WriteString(string(decodedBytes))
}
