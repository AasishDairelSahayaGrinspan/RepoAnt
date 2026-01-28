package cmd

import (
	"fmt"

	"gitsafe-rm/internal/config"
	"gitsafe-rm/internal/github"
	"gitsafe-rm/internal/ui"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check login status and token permissions",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintBanner()
		ui.PrintHeader("Token Status")

		token, err := config.LoadToken()
		if err != nil {
			ui.PrintError("Not logged in")
			ui.PrintHint("Run 'safe-rm login' first")
			fmt.Println()
			return
		}

		ui.PrintSuccess("Token found")
		fmt.Println()

		ui.PrintLoading("Verifying token with GitHub...")
		fmt.Println()

		client := github.NewClient(token)
		hasDeleteScope, scopes, err := client.CheckTokenScopes()
		if err != nil {
			ui.PrintError("Failed to verify token: %v", err)
			ui.PrintHint("Your token may be expired or invalid")
			fmt.Println()
			return
		}

		ui.PrintInfo("Token scopes: %s", scopes)
		fmt.Println()

		if hasDeleteScope {
			ui.Success.Println("  ╔════════════════════════════════════════════════════════╗")
			ui.Success.Println("  ║  ✅ TOKEN HAS 'delete_repo' SCOPE                      ║")
			ui.Success.Println("  ║  You can delete repositories!                          ║")
			ui.Success.Println("  ╚════════════════════════════════════════════════════════╝")
		} else {
			ui.Error.Println("  ╔════════════════════════════════════════════════════════╗")
			ui.Error.Println("  ║  ❌ TOKEN MISSING 'delete_repo' SCOPE                  ║")
			ui.Error.Println("  ║  You CANNOT delete repositories with this token!       ║")
			ui.Error.Println("  ╚════════════════════════════════════════════════════════╝")
			fmt.Println()
			ui.PrintHint("Create a new token with 'delete_repo' scope at:")
			ui.Cyan.Println("  https://github.com/settings/tokens")
			fmt.Println()
			ui.PrintHint("Then run: ./safe-rm login")
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
