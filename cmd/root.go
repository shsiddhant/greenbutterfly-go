package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "greenbutterfly-go",
	Short: "happiness is a green butterfly...",
	Long: `happiness is a green butterfly...
...and she has flown away

A CLI for turning my scattered scripts into commands.`,
	SilenceUsage: true,
	Run:          VersionFunc,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().BoolVar(&showVersion, "version", false, "See version.")
}
