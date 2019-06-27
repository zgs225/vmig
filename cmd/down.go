package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var (
	downCmdVersion string
	downCmdEnv     string
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Rollback migrations",
	Long:  "Rollback migrations",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		app := GetVmigApp()
		v := app.Config.Current.Version
		e := app.Config.Current.Env
		if len(downCmdVersion) > 0 {
			v = downCmdVersion
		}
		if len(downCmdEnv) > 0 {
			e = downCmdEnv
		}
		var err error
		if len(args) == 0 {
			err = app.Down(e, v)
		} else {
			n, err2 := strconv.ParseInt(args[0], 10, 64)
			if err2 != nil {
				panic(err2)
			}
			err = app.DownN(e, v, int(n))
		}
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	downCmd.Flags().StringVarP(&downCmdVersion, "version", "v", "", "Version of migration rollback")
	downCmd.Flags().StringVarP(&downCmdEnv, "env", "e", "", "Environment of migration rollback")
}
