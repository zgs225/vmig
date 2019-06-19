package cmd

import (
	"github.com/spf13/cobra"
)

var (
	newCmdVersion string
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a migration file in default version directory.",
	Long:  "Create a migration file in default version directory, default directory is version directory.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app := GetVmigApp()
		title := args[0]
		version := newCmdVersion
		if version == "" {
			version = app.Config.Current.Version
		}
		if err := app.New(title, version); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&newCmdVersion, "version", "", "", "Version of migration files belongs to.")
}
