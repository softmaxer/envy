package vault

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/softmaxer/envy/pkg/styles"
)

func List(packsDir string) {
	files, err := os.ReadDir(packsDir)
	if err != nil {
		log.Fatal(styles.ErrorText().Render(err.Error()))
		return
	}
	for _, file := range files {
		fmt.Println(styles.ListItem().Render(strings.Split(file.Name(), ".")[0]))
	}
}
