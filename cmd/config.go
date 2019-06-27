package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config settings",
	Long: `Config settings in .vmig.yaml. Support dots form syntax, key is case insensitive.
For example: Config current default version can execute vmig config current.version v1.1.0`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		app := GetVmigApp()
		app.Logger.Debugf("Config %s to value %s", args[0], args[1])
		if err := app.Config.Set(args[0], args[1]); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
