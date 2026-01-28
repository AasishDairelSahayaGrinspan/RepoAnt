package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gitsafe-rm/internal/protected"
	"gitsafe-rm/internal/ui"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var protectCmd = &cobra.Command{
	Use:   "protect",
	Short: "Manage protected repositories",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintBanner()
		ui.PrintHeader("Protected Repositories")
		ui.PrintInfo("Protected repos cannot be deleted with safe-rm")
		fmt.Println()

		protectedRepos, _ := protected.LoadProtectedRepos()

		if len(protectedRepos) == 0 {
			ui.PrintWarning("No protected repositories configured")
			ui.PrintHint("Use 'safe-rm protect add <owner/repo>' to add one")
		} else {
			ui.PrintCount(len(protectedRepos), "protected repository", "protected repositories")
			for repo := range protectedRepos {
				ui.Muted.Print("  🔒 ")
				ui.White.Println(repo)
			}
		}
		fmt.Println()
	},
}

var protectAddCmd = &cobra.Command{
	Use:   "add [owner/repo]",
	Short: "Add a repository to the protected list",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ui.PrintBanner()

		var repoName string
		if len(args) > 0 {
			repoName = args[0]
		} else {
			prompt := &survey.Input{
				Message: "Enter repository name (owner/repo):",
			}
			if err := survey.AskOne(prompt, &repoName); err != nil {
				return err
			}
		}

		repoName = strings.TrimSpace(repoName)
		if repoName == "" || !strings.Contains(repoName, "/") {
			ui.PrintError("Invalid repository name. Use format: owner/repo")
			return fmt.Errorf("invalid repo name")
		}

		if err := addToProtectedList(repoName); err != nil {
			ui.PrintError("Failed to add repository: %v", err)
			return err
		}

		fmt.Println()
		ui.PrintSuccess("Added '%s' to protected list", repoName)
		ui.PrintHint("This repository will no longer appear in delete selection")
		fmt.Println()
		return nil
	},
}

var protectRemoveCmd = &cobra.Command{
	Use:   "remove [owner/repo]",
	Short: "Remove a repository from the protected list",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ui.PrintBanner()

		protectedRepos, _ := protected.LoadProtectedRepos()
		if len(protectedRepos) == 0 {
			ui.PrintWarning("No protected repositories to remove")
			return nil
		}

		var repoName string
		if len(args) > 0 {
			repoName = args[0]
		} else {
			var options []string
			for repo := range protectedRepos {
				options = append(options, repo)
			}
			prompt := &survey.Select{
				Message: "Select repository to unprotect:",
				Options: options,
			}
			if err := survey.AskOne(prompt, &repoName); err != nil {
				ui.PrintInfo("Operation cancelled")
				return nil
			}
		}

		if err := removeFromProtectedList(repoName); err != nil {
			ui.PrintError("Failed to remove repository: %v", err)
			return err
		}

		fmt.Println()
		ui.PrintSuccess("Removed '%s' from protected list", repoName)
		ui.PrintWarning("This repository can now be deleted with safe-rm")
		fmt.Println()
		return nil
	},
}

func getProtectedPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".protected-repos")
}

func addToProtectedList(repoName string) error {
	path := getProtectedPath()

	// Check if already exists
	existing, _ := protected.LoadProtectedRepos()
	if existing[repoName] {
		return fmt.Errorf("repository already protected")
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(repoName + "\n")
	return err
}

func removeFromProtectedList(repoName string) error {
	path := getProtectedPath()

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != repoName {
			lines = append(lines, line)
		}
	}
	file.Close()

	return os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0644)
}

func init() {
	protectCmd.AddCommand(protectAddCmd)
	protectCmd.AddCommand(protectRemoveCmd)
	rootCmd.AddCommand(protectCmd)
}
