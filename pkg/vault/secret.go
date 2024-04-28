package vault

import (
	"strings"

	"github.com/google/uuid"
)

func createSecret() string {
	var secret strings.Builder
	secret.WriteString("envy_")
	secret.WriteString(uuid.New().String())
	return secret.String()
}
