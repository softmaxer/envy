package vault

const (
	secretsFileName string = "envy.secret"
)

func NewVault(root string) {
	createSecretAndPacks(root)
	vaultSecret := createSecret()
	writeSecret(root, vaultSecret)
}

func Pack(filePath string) {
}
