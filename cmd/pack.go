package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "encrypt your env file into a key value store",
	Long: `Usage: 
  envy pack OR envy pack -f <filename>.
  If no filename is provided, envy will look for a .env file present in the current directory.
  If none found, it'll quit gracefully.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pack called")
		filePath, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
		}
		fmt.Printf("Found file path: %s\n", filePath)
	},
}

func init() {
	packCmd.Flags().
		StringP("file", "f", ".env", "Specify a file path for an env file or leave empty if an env file is present in the current directory.")
	rootCmd.AddCommand(packCmd)
}
