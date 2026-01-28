package cmd

import (
	"os"

	"gitsafe-rm/internal/ui"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "reposweep",
	Short: "Safely delete GitHub repositories",
	Long:  "A CLI tool for safely deleting GitHub repositories with interactive selection and protection lists.",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintBanner()
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		ui.PrintError("%v", err)
		os.Exit(1)
	}
}
