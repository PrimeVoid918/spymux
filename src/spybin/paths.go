package spybin

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	systemApplicationsPath = "/usr/share/applications"
	// localApplicationsPath  = "/usr/local/share/applications" //! causes -> err: failed scanning /usr/local/share/applications: lstat /usr/local/share/applications: no such file or directory
)

func applicationSearchPaths() ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to locate home directory: %w", err)
	}

	return []string{
		systemApplicationsPath,
		// localApplicationsPath,
		filepath.Join(home, ".local", "share", "applications"),
	}, nil
}
