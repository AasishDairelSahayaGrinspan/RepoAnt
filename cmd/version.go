package cmd

import (
	"fmt"

	"repoant/internal/ui"

	"github.com/spf13/cobra"
)

var Version = "0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintBanner()
		ui.Primary.Printf("  Version: ")
		ui.White.Printf("%s\n", Version)
		ui.Muted.Printf("  by @aasishdairel\n")
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
