package spybin

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	icons "spymux/src/icons"
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
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}
				if !strings.HasSuffix(info.Name(), ".desktop") {
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
				icon := ResolveAppIcons(name)
				displayName := fmt.Sprintf("%s %s", icon, name)
				if !seen[displayName] {
					seen[displayName] = true
					results = append(
						results,
						AppEntry{
							Name:        name,
							DisplayName: displayName,
							ExecCmd:     execCmd,
						},
					)
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

func ResolveAppIcons(name string) string {
	appIcons := icons.New().Apps
	low := strings.ToLower(name)
	switch {
	case strings.Contains(low, "browser"), strings.Contains(low, "zen"), strings.Contains(low, "chrome"), strings.Contains(low, "tor"):
		return appIcons.Browser
	case strings.Contains(low, "terminal"), strings.Contains(low, "kitty"):
		return appIcons.Terminal
	case strings.Contains(low, "steam"), strings.Contains(low, "game"):
		return appIcons.Game
	case strings.Contains(low, "code"), strings.Contains(low, "obsidian"):
		return appIcons.CodeEditor
	case strings.Contains(low, "torrent"), strings.Contains(low, "qbittorrent"):
		return appIcons.Torrent
	case strings.Contains(low, "bluetooth"), strings.Contains(low, "overskride"):
		return appIcons.Bluetooth
	case strings.Contains(low, "file"), strings.Contains(low, "ncdu"):
		return appIcons.FileManager
	case strings.Contains(low, "audio"), strings.Contains(low, "mixer"), strings.Contains(low, "wiremix"):
		return appIcons.Audio
	default:
		return appIcons.Default
	}
}
