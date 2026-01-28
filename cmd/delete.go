package cmd

import (
	"fmt"
	"os"

	"repoant/internal/config"
	"repoant/internal/github"
	"repoant/internal/protected"
	"repoant/internal/ui"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var multiDelete bool

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Interactively select and delete a repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		ui.PrintBanner()

		// Safety warning banner
		ui.Warning.Println("  ╔════════════════════════════════════════════════════════╗")
		ui.Warning.Println("  ║  ⚠️  WARNING: This will PERMANENTLY delete repository   ║")
		ui.Warning.Println("  ║  This action CANNOT be undone!                         ║")
		ui.Warning.Println("  ╚════════════════════════════════════════════════════════╝")
		fmt.Println()

		token, err := config.LoadToken()
		if err != nil {
			ui.PrintError("Not logged in")
			ui.PrintHint("Run 'repoant login' first")
			fmt.Println()
			return err
		}

		client := github.NewClient(token)

		// Verify token has delete_repo scope
		ui.PrintLoading("Verifying token permissions...")
		hasDeleteScope, scopes, err := client.CheckTokenScopes()
		if err != nil {
			ui.PrintError("Failed to verify token: %v", err)
			return err
		}

		if !hasDeleteScope {
			fmt.Println()
			ui.Error.Println("  ╔════════════════════════════════════════════════════════╗")
			ui.Error.Println("  ║  ❌ TOKEN MISSING 'delete_repo' SCOPE                   ║")
			ui.Error.Println("  ╚════════════════════════════════════════════════════════╝")
			fmt.Println()
			ui.PrintError("Your token does not have permission to delete repositories")
			ui.Muted.Printf("  Current scopes: %s\n", scopes)
			fmt.Println()
			ui.PrintHint("Create a new token with 'delete_repo' scope at:")
			ui.Cyan.Println("  https://github.com/settings/tokens")
			fmt.Println()
			ui.PrintHint("Then run: repoant login")
			fmt.Println()
			return fmt.Errorf("token missing delete_repo scope")
		}

		ui.PrintSuccess("Token verified - has delete_repo permission")
		fmt.Println()

		ui.PrintLoading("Fetching repositories from GitHub...")
		fmt.Println()

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

		protectedRepos, err := protected.LoadProtectedRepos()
		if err != nil {
			ui.PrintError("Failed to load protected repos: %v", err)
			return err
		}

		var repoNames []string
		repoMap := make(map[string]github.Repository)
		protectedCount := 0
		for _, repo := range repos {
			if protectedRepos[repo.FullName] {
				protectedCount++
			} else {
				repoNames = append(repoNames, repo.FullName)
				repoMap[repo.FullName] = repo
			}
		}

		if len(repoNames) == 0 {
			ui.PrintWarning("No deletable repositories found (all are protected)")
			fmt.Println()
			return nil
		}

		ui.PrintHeader("Delete Repository")
		ui.PrintInfo("Total repositories: %d", len(repos))
		if protectedCount > 0 {
			ui.PrintInfo("Protected (hidden): %d", protectedCount)
		}
		ui.PrintInfo("Available to delete: %d", len(repoNames))
		fmt.Println()

		// Multi-select or single-select based on flag
		if multiDelete {
			return handleMultiDelete(repoNames, repoMap, client)
		}
		return handleSingleDelete(repoNames, repoMap, client)
	},
}

func handleSingleDelete(repoNames []string, repoMap map[string]github.Repository, client *github.Client) error {
	ui.Info.Println("  📋 Use ↑↓ arrow keys to navigate, Enter to select")
	ui.Muted.Println("     Press Ctrl+C to cancel at any time")
	fmt.Println()

	var selected string
	prompt := &survey.Select{
		Message: "Select ONE repository to delete:",
		Options: repoNames,
		PageSize: 15,
	}

	// Use custom stdio for better terminal handling
	if err := survey.AskOne(prompt, &selected, survey.WithStdio(os.Stdin, os.Stdout, os.Stderr)); err != nil {
		fmt.Println()
		ui.PrintInfo("Operation cancelled - no repositories deleted")
		fmt.Println()
		return nil
	}

	repo := repoMap[selected]

	// Show what will be deleted
	fmt.Println()
	ui.PrintHeader("Deletion Summary")
	ui.Warning.Println("  ┌─────────────────────────────────────┐")
	ui.Warning.Printf("  │  Repository: %-22s │\n", repo.Name)
	ui.Warning.Printf("  │  Owner:      %-22s │\n", repo.Owner)
	ui.Warning.Printf("  │  Full name:  %-22s │\n", repo.FullName)
	ui.Warning.Println("  │  Count:      1 repository           │")
	ui.Warning.Println("  └─────────────────────────────────────┘")
	fmt.Println()


	// Confirmation prompt
	var confirm bool
	confirmPrompt := &survey.Confirm{
		Message: fmt.Sprintf("Are you SURE you want to delete '%s'? This CANNOT be undone!", repo.FullName),
		Default: false,
	}
	if err := survey.AskOne(confirmPrompt, &confirm); err != nil || !confirm {
		fmt.Println()
		ui.PrintSuccess("Deletion cancelled - repository is safe")
		fmt.Println()
		return nil
	}

	// Final warning
	fmt.Println()
	ui.PrintDeleting(repo.FullName)

	if err := client.DeleteRepository(repo.Owner, repo.Name); err != nil {
		ui.PrintError("Failed to delete repository: %v", err)
		return err
	}

	ui.PrintDeleted(repo.FullName)
	fmt.Println()
	return nil
}

func handleMultiDelete(repoNames []string, repoMap map[string]github.Repository, client *github.Client) error {
	ui.Info.Println("  📋 Use ↑↓ arrow keys to navigate")
	ui.Info.Println("  📋 Press SPACE to select/deselect repositories")
	ui.Info.Println("  📋 Press Enter when done selecting")
	ui.Muted.Println("     Press Ctrl+C to cancel at any time")
	fmt.Println()

	var selected []string
	prompt := &survey.MultiSelect{
		Message: "Select repositories to delete (SPACE to toggle, ENTER to confirm):",
		Options: repoNames,
		PageSize: 15,
	}

	if err := survey.AskOne(prompt, &selected, survey.WithStdio(os.Stdin, os.Stdout, os.Stderr)); err != nil {
		fmt.Println()
		ui.PrintInfo("Operation cancelled - no repositories deleted")
		fmt.Println()
		return nil
	}

	if len(selected) == 0 {
		fmt.Println()
		ui.PrintWarning("No repositories selected")
		ui.PrintInfo("Use SPACE key to select repositories before pressing Enter")
		fmt.Println()
		return nil
	}

	// Show what will be deleted
	fmt.Println()
	ui.PrintHeader("Deletion Summary")
	ui.Error.Printf("  ⚠️  YOU ARE ABOUT TO DELETE %d REPOSITORY(IES):\n", len(selected))
	fmt.Println()
	for i, name := range selected {
		ui.Warning.Printf("  %2d. 🗑️  %s\n", i+1, name)
	}
	fmt.Println()
	ui.Error.Println("  ⚠️  THIS ACTION CANNOT BE UNDONE!")
	fmt.Println()


	// Double confirmation for multi-delete
	var confirm bool
	confirmPrompt := &survey.Confirm{
		Message: fmt.Sprintf("Are you SURE you want to delete these %d repositories?", len(selected)),
		Default: false,
	}
	if err := survey.AskOne(confirmPrompt, &confirm); err != nil || !confirm {
		fmt.Println()
		ui.PrintSuccess("Deletion cancelled - all repositories are safe")
		fmt.Println()
		return nil
	}

	// Type confirmation for safety
	var typeConfirm string
	typePrompt := &survey.Input{
		Message: fmt.Sprintf("Type 'DELETE %d' to confirm:", len(selected)),
	}
	if err := survey.AskOne(typePrompt, &typeConfirm); err != nil {
		fmt.Println()
		ui.PrintSuccess("Deletion cancelled - all repositories are safe")
		fmt.Println()
		return nil
	}

	expectedConfirm := fmt.Sprintf("DELETE %d", len(selected))
	if typeConfirm != expectedConfirm {
		fmt.Println()
		ui.PrintError("Confirmation text did not match")
		ui.PrintSuccess("Deletion cancelled - all repositories are safe")
		fmt.Println()
		return nil
	}

	// Delete repositories
	fmt.Println()
	ui.PrintHeader("Deleting Repositories")

	successCount := 0
	failCount := 0

	for _, name := range selected {
		repo := repoMap[name]
		ui.Warning.Printf("  🗑️  Deleting %s... ", name)

		if err := client.DeleteRepository(repo.Owner, repo.Name); err != nil {
			ui.Error.Println("FAILED")
			ui.PrintError("  Error: %v", err)
			failCount++
		} else {
			ui.Success.Println("DONE")
			successCount++
		}
	}

	// Summary
	fmt.Println()
	ui.PrintHeader("Deletion Complete")
	if successCount > 0 {
		ui.PrintSuccess("Successfully deleted: %d repositories", successCount)
	}
	if failCount > 0 {
		ui.PrintError("Failed to delete: %d repositories", failCount)
	}
	fmt.Println()

	return nil
}

func init() {
	deleteCmd.Flags().BoolVarP(&multiDelete, "multi", "m", false, "Select multiple repositories to delete at once")
	rootCmd.AddCommand(deleteCmd)
}
