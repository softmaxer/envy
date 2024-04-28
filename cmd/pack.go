package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/softmaxer/envy/pkg/vault"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "encrypt your env file into a key value store",
	Long: `Usage: 
  envy pack OR envy pack -f <filename>.
  If no filename is provided, envy will look for a .env file present in the current directory.
  If none found, it'll quit gracefully.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envPath, err := cmd.Flags().GetString("file")
		var projectName string = vault.GetProjectName()
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
		}
		if len(args) > 0 {
			projectName = args[0]
		}
		envFD, err := os.Open(envPath)
		if err != nil {
			log.Fatalf("Error opening env: %s\n", err.Error())
		}
		vault.Pack(projectName, envFD)
	},
}

func init() {
	packCmd.Flags().
		StringP("file", "f", ".env", "Specify a file path for an env file or leave empty if an env file is present in the current directory.")
	rootCmd.AddCommand(packCmd)
}
