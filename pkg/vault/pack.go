package vault

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/softmaxer/envy/pkg/styles"
)

func Pack(project string, envFile io.Reader) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(styles.ErrorText().Render("Cannot access home folder: ", err.Error()))
	}
	packExt := fmt.Sprintf("%s.pack", project)
	packName := filepath.Join(homeDir, ".envy", "packs", packExt)
	packFD, err := os.Create(packName)
	if err != nil {
		log.Fatal(styles.ErrorText().Render("Error creating file: ", err.Error()))
	}
	defer packFD.Close()
	keyName := filepath.Join(homeDir, ".envy", secretsFileName)
	keyFD, err := os.OpenFile(keyName, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(styles.ErrorText().Render("Error creating secret key: ", err.Error()))
	}
	defer keyFD.Close()
	hash, err := encrypt(keyFD, envFile)
	if err != nil {
		log.Fatal(styles.ErrorText().Render("Error encoding the file: ", err.Error()))
	}
	packFD.Write(hash)
	log.Print(styles.SuccessText().Render("Successfully packed ", project))
}
