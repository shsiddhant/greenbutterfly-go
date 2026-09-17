package cmd

import (
	"os"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"

	"github.com/shsiddhant/greenbutterfly-go/ui"
)

var quoteText = ui.QuoteStyle.Render(`
"happiness is a green butterfly...
...and she has flown away"`)

var longText = quoteText + "\n\nA CLI for turning my scattered scripts into commands."

var rootCmd = &cobra.Command{
	Use:          "greenbutterfly-go",
	Short:        "happiness is a green butterfly...",
	Long:         longText,
	SilenceUsage: true,
	Run:          VersionFunc,
}

func Execute() {
	cc.Init(&cc.Config{
		RootCmd:         rootCmd,
		Headings:        cc.HiYellow + cc.Bold,
		Commands:        cc.HiCyan + cc.Bold,
		Flags:           cc.HiGreen + cc.Bold,
		ExecName:        cc.Bold + cc.HiMagenta,
		Aliases:         cc.Italic,
		NoExtraNewlines: true,
		NoBottomNewline: true,
	})
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().BoolVar(&showVersion, "version", false, "See version.")
}
