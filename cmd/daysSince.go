/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"unicode/utf8"

	"github.com/spf13/cobra"

	"github.com/shsiddhant/greenbutterfly-go/internal"
	"github.com/shsiddhant/greenbutterfly-go/ui"
)

var jsonFilePath string

// daysSinceCmd represents the days-since command.
//
// Aliases: ds, days_since
var daysSinceCmd = &cobra.Command{
	Use:     "days-since JSON_FILE_PATH",
	Aliases: []string{"ds", "days_since"},
	Args:    cobra.ExactArgs(1),
	Short: "View days since for a list of significant dates." +
		" (aliases: ds, days_since)",
	Long: `View days since for a list of significant dates.
Sorted oldest to newest.`,

	Example: `  $ greenbutterfly-go days-since dates.json
   DAYS SINCE ────────────────────────────
   1005 • The Best Day • December 17, 2023`,
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
		for i, day := range daysSince {
			countText := fmt.Sprint(
				ui.DaysCountStyle.Render(fmt.Sprintf("%5d", day.Days)),
			)
			metaText := fmt.Sprintf(
				" • %s • %s\n",
				ui.TitleStyle.Render(day.Title),
				ui.DimStyle.Render(day.Date),
			)
			if i == 0 {
				formattedString := fmt.Sprintf(
					"%5d • %s • %s",
					day.Days,
					day.Title,
					day.Date,
				)
				fmt.Print(
					ui.RenderHeader(
						"Days Since",
						utf8.RuneCountInString(formattedString),
					),
				)
			}
			fmt.Print(countText + metaText)
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
