package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/softmaxer/envy/pkg/styles"
	"github.com/softmaxer/envy/pkg/vault"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "encrypt your env file into an AES encrypted file.",
	Long: `Usage: 
  envy pack OR envy pack -f <filename>.
  This will pack all the environment variables in a given env file with the name of the directory of your project.
  additionally, you can also specify a project name as a positional argument.
  example: envy pack my-cool-project. this will pack your env into a project called "my-cool-project".
  you can also specify a project name along with a filename.
  Note: This will attempt to remove the .env file from the path specified and will throw an error if it doesn't succeed.
  If no filename is provided, envy will look for a .env file present in the current directory.
  If none found, The program will quit with an error.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envPath, err := cmd.Flags().GetString("file")
		var projectName string = vault.GetProjectName()
		if err != nil {
			log.Fatal(styles.ErrorText().Render("Error: ", err.Error()))
		}
		if len(args) > 0 {
			projectName = args[0]
		}
		envFD, err := os.Open(envPath)
		if err != nil {
			log.Fatal(styles.ErrorText().Render("Error opening env: ", err.Error()))
		}

		defer envFD.Close()
		vault.Pack(projectName, envFD)
		err = os.Remove(envPath)
		if err != nil {
			log.Fatal(
				styles.ErrorText().Render("Error removing .env from file path: ", err.Error()),
			)
		}
	},
}

func init() {
	packCmd.Flags().
		StringP("file", "f", ".env", "Specify a file path for an env file or leave empty if an env file is present in the current directory.")
	rootCmd.AddCommand(packCmd)
}
