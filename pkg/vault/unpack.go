package vault

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/softmaxer/envy/pkg/styles"
)

func Unpack(project string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(styles.ErrorText().Render("Unable to access home: ", err.Error()))
	}
	packExt := fmt.Sprintf("%s.pack", project)
	packName := filepath.Join(homeDir, ".envy", "packs", packExt)
	packFD, err := os.Open(packName)
	if err != nil {
		log.Fatal(styles.ErrorText().Render("Error opening pack: ", err.Error()))
	}
	defer packFD.Close()
	keyName := filepath.Join(homeDir, ".envy", secretsFileName)
	keyFD, err := os.OpenFile(keyName, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(styles.ErrorText().Render("Error getting secret key: ", err.Error()))
	}
	defer keyFD.Close()
	decodedEnv, err := decrypt(keyFD, packFD)
	WriteDecodedPack(decodedEnv)
	log.Print(styles.SuccessText().Render("Unpacked ", project, " into .env!"))
}
