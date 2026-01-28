package cmd

import (
	"fmt"

	"repoant/internal/config"
	"repoant/internal/ui"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Store GitHub personal access token",
	RunE: func(cmd *cobra.Command, args []string) error {
		ui.PrintBanner()
		ui.PrintHeader("GitHub Authentication")
		ui.PrintInfo("Token requires scopes: repo, delete_repo")
		ui.PrintHint("Create token at: https://github.com/settings/tokens")

		fmt.Println()
		var token string
		prompt := &survey.Password{
			Message: "Enter your GitHub Personal Access Token:",
		}
		if err := survey.AskOne(prompt, &token); err != nil {
			ui.PrintError("Failed to read token: %v", err)
			return err
		}

		if token == "" {
			ui.PrintError("Token cannot be empty")
			return fmt.Errorf("token cannot be empty")
		}

		if err := config.SaveToken(token); err != nil {
			ui.PrintError("Failed to save token: %v", err)
			return err
		}

		fmt.Println()
		ui.PrintSuccess("Token saved successfully!")
		ui.PrintHint("Run 'repoant list' to see your repositories")
		fmt.Println()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
