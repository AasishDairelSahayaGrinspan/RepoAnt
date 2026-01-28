package cmd

import (
	"fmt"

	"repoant/internal/config"
	"repoant/internal/github"
	"repoant/internal/ui"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your GitHub repositories",
	RunE: func(cmd *cobra.Command, args []string) error {
		ui.PrintBanner()

		token, err := config.LoadToken()
		if err != nil {
			ui.PrintError("Not logged in")
			ui.PrintHint("Run 'safe-rm login' first")
			fmt.Println()
			return err
		}

		ui.PrintLoading("Fetching repositories from GitHub...")
		fmt.Println()

		client := github.NewClient(token)
		repos, err := client.ListRepositories()
		if err != nil {
			ui.PrintError("Failed to fetch repositories: %v", err)
			return err
		}

		if len(repos) == 0 {
			ui.PrintWarning("No repositories found")
			fmt.Println()
			return nil
		}

		ui.PrintHeader("Your Repositories")
		ui.PrintCount(len(repos), "repository", "repositories")

		for _, repo := range repos {
			ui.PrintRepoSimple(repo.FullName)
		}

		fmt.Println()
		ui.PrintHint("Run 'repoant delete' to interactively delete a repository")
		fmt.Println()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
