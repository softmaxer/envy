package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/softmaxer/envy/pkg/styles"
	"github.com/softmaxer/envy/pkg/vault"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all encrypted env files",
	Long: `Usage: envy list, lists all 
  the available packs so far to browse through your project variables 
  if you ever forget a name of a project that yout previously packed.`,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		vaultDir := filepath.Join(homeDir, ".envy")
		packsDir := filepath.Join(vaultDir, "packs")
		if err != nil {
			log.Fatal(styles.ErrorText().Render(err.Error()))
		}
		vault.List(packsDir)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
