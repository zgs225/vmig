package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var (
	upCmdVersion string
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply all or given N up migration files.",
	Long:  "Apply all or given N up migration files.",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		app := GetVmigApp()
		v := app.Config.Current.Version
		if len(upCmdVersion) > 0 {
			v = upCmdVersion
		}
		var err error
		if len(args) == 0 {
			err = app.Up(v)
		} else {
			n, err2 := strconv.ParseInt(args[0], 10, 64)
			if err2 != nil {
				panic(err2)
			}
			err = app.UpN(v, int(n))
		}
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	upCmd.Flags().StringVarP(&upCmdVersion, "version", "", "", "Version of up")
}
