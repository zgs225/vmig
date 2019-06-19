package cmd

import (
	"github.com/spf13/cobra"
)

var (
	newCmdDir string
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
		dir := newCmdDir
		if dir == "" {
			dir = app.Config.Current.Version
		}
		if err := app.New(title, dir); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&newCmdDir, "dir", "d", "", "Directory for create migration files.")
}
