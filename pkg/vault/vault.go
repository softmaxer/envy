package vault

import (
	"log"

	"github.com/softmaxer/envy/pkg/styles"
)

const (
	secretsFileName string = "envy.secret"
)

func NewVault(root string) {
	createSecretAndPacks(root)
	vaultSecret := createSecret()
	writeSecret(root, vaultSecret)
	log.Print(styles.SuccessText().Render("Created a new vault!"))
}
