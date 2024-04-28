package vault

import (
	"crypto/rand"
	"log"

	"github.com/softmaxer/envy/pkg/styles"
)

func createSecret() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal(styles.ErrorText().Render("Error reading into key bytes: ", err.Error()))
	}
	return string(key)
}
