package cmd

import (
	"os"

	"repoant/internal/ui"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "repoant",
	Short: "Delete and mass delete GitHub repositories with a single click",
	Long:  "🐜 RepoAnt - A CLI tool for deleting GitHub repositories. Select one or multiple repos and delete them with a single confirmation.",
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
