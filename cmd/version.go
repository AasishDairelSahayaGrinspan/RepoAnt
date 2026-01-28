package cmd

import (
	"fmt"

	"gitsafe-rm/internal/ui"

	"github.com/spf13/cobra"
)

var Version = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintBanner()
		ui.Primary.Printf("  Version: ")
		ui.White.Printf("%s\n", Version)
		ui.Muted.Printf("  Built with ❤️  using Go + Cobra\n")
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
