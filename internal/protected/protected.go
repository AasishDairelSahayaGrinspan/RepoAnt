package protected

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

const protectedFileName = ".protected-repos"

func getProtectedPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, protectedFileName), nil
}

func LoadProtectedRepos() (map[string]bool, error) {
	protected := make(map[string]bool)

	path, err := getProtectedPath()
	if err != nil {
		return protected, nil
	}

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return protected, nil
		}
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			protected[line] = true
		}
	}

	return protected, scanner.Err()
}
