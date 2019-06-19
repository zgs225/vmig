package cmd

import (
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init vmig environment",
	Long:  `Initial vmig configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		app := GetVmigApp()
		if err := app.Init(); err != nil {
			panic(err)
		}
		app.Logger.Info("Initialized")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
