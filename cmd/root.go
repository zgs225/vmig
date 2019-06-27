package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zgs225/vmig/core"
)

var (
	cfgFile string
	Verbose bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vmig",
	Short: "Versiond && multiple environments golang-migrate",
	Long: `A wrapper for golang-migrate that let it support version managed migrations and
multiple environments.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		app, err := core.NewApp("")
		if err != nil {
			panic(err)
		}
		SetVmigApp(app)
		app.SetVerbose(Verbose)
		if err := app.LoadConfigFromViper(viper.GetViper()); err != nil {
			panic(err)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		app := GetVmigApp()
		if app.Config.IsDirty() {
			if err := app.DumpConfigByViper(viper.GetViper()); err != nil {
				panic(err)
			}
			app.Logger.WithField("file", viper.GetViper().ConfigFileUsed()).Debug("Config dumped into file.")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PWD/.vmig.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "", false, "Output debug message.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find pwd directory.
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		// Search config in pwd directory with name ".vmig" (without extension).
		viper.AddConfigPath(dir)
		viper.SetConfigName(core.DEFAULT_CONFIG_FILE)
	}

	// viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	viper.ReadInConfig()
}
