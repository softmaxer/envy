package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "envy",
	Short: "An environment variable manager for your git repositories",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
Envy is a simple CLI tool to manage your .env files in your git repos.
It uses an AES encryption to store all your .env files in a vault.
With a simple command, they're always at your disposal to either pack / unpack them
into your current projects.
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
