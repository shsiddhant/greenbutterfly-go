package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Version = "v0.0.0" // Fallback

var showVersion bool

var VersionFunc = func(cmd *cobra.Command, args []string) {
	if showVersion {
		fmt.Printf("greenbutterfly-go version %s\n", strings.TrimLeft(Version, "v"))
		fmt.Println("Copyright (c) 2026 Siddhant Sharma")
		os.Exit(0)
	}
	cmd.Help()
}
