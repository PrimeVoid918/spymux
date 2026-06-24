package utils

import (
	"errors"
	"fmt"
	"os"
	"path"
)

func HomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to locate home directory: %w", err)
	}

	return path.Clean(homeDir) + "/", nil
}

func WalCacheDir() (string, error) {
	cachePath := ".cache/wal/colors.json"

	homePath, pathErr := HomeDir()
	if pathErr != nil {
		return "", pathErr
	}

	fullPath := homePath + cachePath
	info, err := os.Stat(fullPath)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", errors.New("wal colors.json cache file is missing")
		}
		return "", fmt.Errorf("failed checking cache path %s: %w", fullPath, err)
	}

	if info.IsDir() {
		return "", fmt.Errorf("expected a file but found a directory at: %s", fullPath)
	}

	return path.Clean(fullPath), nil
}
