# 🐜 RepoAnt

A CLI tool for deleting GitHub repositories - select one or multiple repos and delete them with a single confirmation.

## Features

- 🎨 **Modern UI**: Sleek terminal interface with gradient colors
- 🔍 **Interactive selection**: Browse and select repositories using arrow keys
- 🔒 **Protected repos**: Configure repositories that can never be deleted
- 🔐 **Secure token storage**: GitHub PAT stored locally with restricted permissions
- ⚠️ **Safety warnings**: Clear warnings and confirmation prompts before deletion
- 📦 **Mass delete**: Delete multiple repositories at once with extra safeguards

## Installation

RepoAnt works on **Windows**, **Linux**, and **macOS**.

### Homebrew (macOS / Linux)

```bash
brew install AasishDairelSahayaGrinspan/repoant/repoant
```

Alternative (two-step):

```bash
brew tap AasishDairelSahayaGrinspan/repoant
brew install repoant
```

### macOS / Linux

```bash
# Build from source
go mod tidy
go build -o repoant .

# Install globally (optional)
sudo mv repoant /usr/local/bin/
```

### Windows

```bash
# Build from source
go mod tidy
go build -o repoant.exe .

# Add to PATH (optional)
# Move repoant.exe to a directory in your PATH
```

### Pre-built Binaries

Download pre-built binaries from the releases page.

## Usage

### Login

Store your GitHub Personal Access Token (requires `repo` and `delete_repo` scopes):

```bash
repoant login
```

### List Repositories

View all your GitHub repositories:

```bash
repoant list
```

### Delete Repository (Single)

Interactively select and delete ONE repository:

```bash
repoant delete
```

**Navigation:**
- ↑↓ Arrow keys to navigate
- Enter to select
- Ctrl+C to cancel

### Delete Multiple Repositories

Select and delete multiple repositories at once:

```bash
repoant delete --multi
# or
repoant delete -m
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
repoant protect
```

Add a repository to the protected list:
```bash
repoant protect add owner/repo
```

Remove a repository from the protected list:
```bash
repoant protect remove owner/repo
```

### Version

Check the CLI version:
```bash
repoant version
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

The token is stored at `~/.repoant-token` with `0600` permissions (readable only by you).

## Project Structure

```
repoant/
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

## Author

by @aasishdairel

## License

MIT

## Cross-Platform Compatibility

RepoAnt is built with Go and works seamlessly on:

- **macOS** (Intel and Apple Silicon)
- **Linux** (x86_64, ARM64)
- **Windows** (x86_64)

The application uses cross-platform libraries for:
- Terminal UI interactions
- File system operations
- Color output
- HTTP requests

All features work identically across platforms.
