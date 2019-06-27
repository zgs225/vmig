package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	Major = 1
	Minor = 0
	Patch = 3
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show vmig version",
	Long:  "Show vmig version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version %d.%d.%d\n", Major, Minor, Patch)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
