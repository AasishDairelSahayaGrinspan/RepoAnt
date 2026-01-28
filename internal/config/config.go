package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".reposweep-token"

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
	return string(data), nil
}
