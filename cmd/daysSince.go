/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/shsiddhant/greenbutterfly-go/internal"
)

var jsonFilePath string

// daysSinceCmd represents the days-since command.
//
// Aliases: ds, days_since
var daysSinceCmd = &cobra.Command{
	Use:     "days-since [JSON_FILE_PATH]",
	Aliases: []string{"ds", "days_since"},
	Short: "View days since for a list of significant dates." +
		" (aliases: ds, days_since)",
	Long: `View days since for a list of significant dates.
Sorted oldest to newest.

Example output:
It has been 1005 days since 'The Best Day' - December 17, 2023.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		jsonFilePath = args[0]
		absPath, err := filepath.Abs(jsonFilePath)
		if err != nil {
			return fmt.Errorf("error resolving absolute file path: %w", err)
		}
		if _, err := os.Stat(absPath); os.IsNotExist(err) {
			return fmt.Errorf("file does not exist: %s", absPath)
		}
		daysSince, err := internal.CalculateDaysSince(absPath)
		if err != nil {
			return err
		}
		for _, day := range daysSince {
			fmt.Printf(
				"It has been %d days since '%s' - %s.\n",
				day.Days,
				day.Title,
				day.Date,
			)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(daysSinceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// daysSinceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// daysSinceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
