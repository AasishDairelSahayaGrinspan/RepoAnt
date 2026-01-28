# RepoSweep

A CLI tool for safely deleting GitHub repositories with interactive selection and protection lists.

## Features

- 🎨 **Colorful TUI**: Beautiful, modern terminal interface with colors and emojis
- 🔍 **Interactive selection**: Browse and select repositories using arrow keys
- 🔒 **Protected repos**: Configure repositories that can never be deleted
- 🔐 **Secure token storage**: GitHub PAT stored locally with restricted permissions
- ⚠️ **Safety warnings**: Clear warnings and confirmation prompts before deletion
- 📦 **Multi-delete**: Option to delete multiple repositories at once with extra safeguards

## Installation

```bash
# Build from source
go mod tidy
go build -o reposweep .

# Move to PATH (optional)
sudo mv reposweep /usr/local/bin/
```

## Usage

### Login

Store your GitHub Personal Access Token (requires `repo` and `delete_repo` scopes):

```bash
./reposweep login
```

### List Repositories

View all your GitHub repositories:

```bash
./reposweep list
```

### Delete Repository (Single)

Interactively select and delete ONE repository:

```bash
./reposweep delete
```

**Navigation:**
- ↑↓ Arrow keys to navigate
- Enter to select
- Ctrl+C to cancel

### Delete Multiple Repositories

Select and delete multiple repositories at once:

```bash
./reposweep delete --multi
# or
./reposweep delete -m
```

**Navigation:**
- ↑↓ Arrow keys to navigate  
- SPACE to toggle selection
- Enter to confirm
- Ctrl+C to cancel

⚠️ Multi-delete requires typing `DELETE <count>` to confirm.

### Manage Protected Repositories

View protected repositories:
```bash
./reposweep protect
```

Add a repository to the protected list:
```bash
./reposweep protect add owner/repo
```

Remove a repository from the protected list:
```bash
./reposweep protect remove owner/repo
```

### Version

Check the CLI version:
```bash
./reposweep version
```

## Protected Repositories

Protected repositories will not appear in the delete selection list. You can manage them with the `protect` command or manually edit `~/.protected-repos`:

```text
# Protected repositories (one per line, format: owner/repo)
myusername/important-repo
myusername/production-app
```

## GitHub Token

The CLI requires a GitHub Personal Access Token with the following scopes:
- `repo` - Full control of private repositories
- `delete_repo` - Delete repositories

Create a token at: https://github.com/settings/tokens

The token is stored at `~/.reposweep-token` with `0600` permissions (readable only by you).

## Project Structure

```
reposweep/
├── main.go                          # Entry point
├── go.mod                           # Go module definition
├── cmd/
│   ├── root.go                      # Root command
│   ├── login.go                     # Login command
│   ├── list.go                      # List command
│   ├── delete.go                    # Delete command
│   ├── protect.go                   # Protect command
│   └── version.go                   # Version command
├── internal/
│   ├── config/
│   │   └── config.go                # Token storage
│   ├── github/
│   │   └── client.go                # GitHub API client
│   ├── protected/
│   │   └── protected.go             # Protected repos handling
│   └── ui/
│       └── ui.go                    # Colorful UI components
└── .protected-repos.example         # Example protected repos file
```

## License

MIT
