package utils

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func HomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to locate home directory: %w", err)
	}

	return path.Clean(homeDir) + "/", nil
}

func WalCachePath(file string) (string, error) {
	homePath, err := HomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(homePath, ".cache", "wal", file)

	info, err := os.Stat(fullPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("wal cache file does not exist: %s", fullPath)
		}
		return "", fmt.Errorf("failed checking wal cache path %s: %w", fullPath, err)
	}

	if info.IsDir() {
		return "", fmt.Errorf("expected a file but found a directory at: %s", fullPath)
	}

	return fullPath, nil
}
