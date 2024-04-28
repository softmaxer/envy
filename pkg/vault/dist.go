package vault

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/softmaxer/envy/pkg/styles"
)

func readLines(envFile io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func createEnvDist(envFile io.Reader) {
	envDist, err := os.Create(".env-dist")
	if err != nil {
		log.Fatal(
			styles.ErrorText().Render("Error creating .env-dist in this directory: ", err.Error()),
		)
	}
	defer envDist.Close()
	lines, err := readLines(envFile)
	fmt.Printf("Lines: %+v\n", lines)
	if err != nil {
		log.Fatal(styles.ErrorText().Render("Error reading lines from env: ", err.Error()))
	}
	for _, line := range lines {
		fmt.Printf("Line: %s\n", line)
		vars := strings.Split(line, "=")
		fmt.Printf("Found: %+v\n", vars)
		if len(vars) < 2 {
			continue
		}
		keyString := fmt.Sprintf("%s=\n", vars[0])
		_, err := envDist.WriteString(keyString)
		if err != nil {
			log.Print(styles.ErrorText().Render("Error writing a variable, skipping"))
			continue
		}
	}
	log.Print(styles.SuccessText().Render("Wrote variable keys to env-dist"))
}
