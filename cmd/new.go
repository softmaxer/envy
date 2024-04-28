/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/softmaxer/envy/pkg/styles"
	"github.com/softmaxer/envy/pkg/vault"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new envy vault",
	Long: `envy new creates a new envy vault in the HOME/.envy folder. 
  It will create a folder .envy if it doesn't already exist.
  envy new also adds a new '.secret' that contains a secret encryption key for all your env files.

use envy new <custom path to envy folder>
To use another path other than the home
.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			vaultDir := filepath.Join(args[0], ".envy")
			vault.NewVault(vaultDir)
		} else {
			homeDir, err := os.UserHomeDir()
			vaultDir := filepath.Join(homeDir, ".envy")
			if err != nil {
				log.Fatal(styles.ErrorText().Render(err.Error()))
			}
			vault.NewVault(vaultDir)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
