package cmd

import (
	"github.com/spf13/cobra"
)

var (
	isDefault bool
)

// createVersionCmd represents the createVersion command
var createVersionCmd = &cobra.Command{
	Use:   "create-version",
	Short: "Create a new version",
	Long:  "Create a new version at working directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app := GetVmigApp()
		v := args[0]
		if err := app.CreateVersion(v, isDefault); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createVersionCmd)

	createVersionCmd.Flags().BoolVarP(&isDefault, "default", "d", false, "Whether is default version")
}
