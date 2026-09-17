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

var nextYear bool

// anniversariesCmd represents the anniversary command
//
// Alias: ann
var anniversariesCmd = &cobra.Command{
	Use:     "anniversaries JSON_FILE_PATH",
	Aliases: []string{"ann"},
	Args:    cobra.ExactArgs(1),
	Short:   "View upcoming anniversaries. (alias: ann)",
	Long: `View upcoming anniversaries for a list of significant dates.
Sorted closest to farthest.`,
	Example: `  $ greenbutterfly-go anniversaries dates.json
   UPCOMING ANNIVERSARIES ─────────────────────────────
  • The Best Day (3rd) — December 17, 2026 • in 90 days`,
	RunE: func(cmd *cobra.Command, args []string) error {
		jsonFilePath = args[0]
		absPath, err := filepath.Abs(jsonFilePath)
		if err != nil {
			return fmt.Errorf("error resolving absolute file path: %w", err)
		}
		if _, err := os.Stat(absPath); os.IsNotExist(err) {
			return fmt.Errorf("file does not exist: %s", absPath)
		}
		anniversaries, err := internal.CalculateAnniversaries(jsonFilePath, nextYear)
		if err != nil {
			return err
		}
		for i, a := range anniversaries {
			if i == 0 {
				formattedString := fmt.Sprintf(
					"• %s (%s) — %s • in %d days",
					a.Title,
					ordinal(a.Number),
					a.Date,
					a.DaysUntil,
				)
				fmt.Print(
					ui.RenderHeader(
						"Upcoming Anniversaries",
						utf8.RuneCountInString(formattedString),
					),
				)
			}
			fmt.Printf(
				"• %s %s — %s • %s\n",
				ui.TitleStyle.Render(a.Title),
				ui.DimStyle.Render(fmt.Sprintf("(%s)", ordinal(a.Number))),
				ui.DimStyle.Render(a.Date),
				ui.DaysCountStyle.Render(fmt.Sprintf("in %d days", a.DaysUntil)),
			)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(anniversariesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// anniversaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// anniversaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	anniversariesCmd.Flags().BoolVarP(
		&nextYear,
		"next-year",
		"n",
		false,
		"Include anniversaries from the next year.",
	)
}

func ordinal(number int) string {
	suffix := "th"
	moduloHundred := number % 100
	moduloTen := number % 10
	switch {
	case moduloHundred >= 11 && moduloHundred <= 13:
	case moduloTen == 1:
		suffix = "st"
	case moduloTen == 2:
		suffix = "nd"
	case moduloTen == 2:
		suffix = "rd"
	}
	return fmt.Sprintf("%d%s", number, suffix)
}
