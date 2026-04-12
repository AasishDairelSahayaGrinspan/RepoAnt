package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const configFileName = ".repoant-token"

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}

func SaveToken(token string) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(token), 0600)
}

func LoadToken() (string, error) {
	path, err := getConfigPath()
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("token not found")
		}
		return "", err
	}
	token := strings.TrimSpace(string(data))
	if token == "" {
		return "", fmt.Errorf("token is empty")
	}
	return token, nil
}
