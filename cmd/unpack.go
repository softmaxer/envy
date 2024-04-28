/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/softmaxer/envy/pkg/vault"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack contents of a project's environment variables.",
	Long: `Use envy unpack to unveil all the environment variables of the project into a .env file.
Additionally, you can also specify a name of a project:
envy unpack my-cool-project
`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string = vault.GetProjectName()
		if len(args) > 0 {
			projectName = args[0]
		}
		vault.Unpack(projectName)
	},
}

func init() {
	rootCmd.AddCommand(unpackCmd)
}
