package spybin

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var ignoredPathFragments = []string{
	"wine",
	"programs",
	"uninstall",
	"chrome apps",
	"google maps",
}
var blacklistedNames = []string{
	"avahi",
	"qt6",
	"qt5",
	"assistant",
	"designer",
	"linguist",
	"openjdk",
	"java",
	"xwayland",
	"console",
	"shell",
	"runtime",
}

var ignoredApps = map[string]struct{}{
	"Base":    {},
	"Math":    {},
	"Draw":    {},
	"Writer":  {},
	"Calc":    {},
	"Impress": {},
}

func ScanDesktopFiles() ([]AppEntry, error) {
	searchPaths, err := applicationSearchPaths()
	if err != nil {
		return nil, err
	}

	var results []AppEntry
	seen := make(map[string]bool)

	for _, path := range searchPaths {
		filepathWalkErr := filepath.Walk(
			path,
			func(
				p string,
				info os.FileInfo,
				err error,
			) error {
				// ignore invalid things
				// if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(), ".desktop") {
				// 	return nil
				// }
				if err != nil {
					fmt.Printf("err is not nil: %s\n", err)
					return err
				}
				if info.IsDir() {
					fmt.Printf("info.IsDir is not nil: %v\n", true)
					return nil
				}
				if !strings.HasSuffix(info.Name(), ".desktop") {
					fmt.Printf("HasSuffix is not nil: %v\n", !strings.HasSuffix(info.Name(), ".desktop"))
					return nil
				}

				// ignore unwanted paths
				if isIgnoredPath(p) {
					return nil
				}

				name, execCmd, valid, scanErr := parseDesktopFile(p)
				if scanErr != nil {
					return scanErr
				}
				if !valid {
					return nil
				}

				// ignore invalid apps
				if _, ignored := ignoredApps[name]; ignored {
					return nil
				}
				if isBlacklistedApp(name) {
					return nil
				}

				// TODO: refactor for icon domain later
				icon := getIcon(name)
				displayName := fmt.Sprintf("%s %s", icon, name)
				if !seen[displayName] {
					seen[displayName] = true
					results = append(results, AppEntry{DisplayName: displayName, ExecCmd: execCmd})
				}
				return nil
			})

		if filepathWalkErr != nil {
			return nil, filepathWalkErr
		}
	}
	return results, nil
}

func (m *Model) filterApps() {
	m.Index = 0
	if m.Query == "" {
		m.Filtered = m.Apps
		return
	}

	var out []AppEntry
	q := strings.ToLower(m.Query)
	for _, app := range m.Apps {
		if strings.Contains(strings.ToLower(app.DisplayName), q) {
			out = append(out, app)
		}
	}
	m.Filtered = out
}

func isIgnoredPath(path string) bool {
	lower := strings.ToLower(path)

	for _, fragment := range ignoredPathFragments {
		if strings.Contains(lower, fragment) {
			return true
		}
	}

	return false
}

func parseDesktopFile(filePath string) (string, string, bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", "", false, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var name, execCmd string
	var inMainSection, skipDisplay bool

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			if line == "[Desktop Entry]" {
				inMainSection = true
			} else {
				inMainSection = false
			}
			continue
		}

		if !inMainSection {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		switch key {
		case "NoDisplay", "Terminal":
			if val == "true" {
				skipDisplay = true
			}
		case "Name":
			if name == "" {
				name = val
			}
		case "Exec":
			if execCmd == "" {
				idx := strings.Index(val, " %")
				if idx != -1 {
					execCmd = val[:idx]
				} else {
					execCmd = val
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", "", false, scanner.Err()
	}

	if skipDisplay || name == "" || execCmd == "" {
		return "", "", false, nil
	}
	return name, execCmd, true, nil
}

func isBlacklistedApp(name string) bool {
	low := strings.ToLower(name)
	for _, b := range blacklistedNames {
		if strings.Contains(low, b) {
			return true
		}
	}
	return false
}
