package config

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type AppConfig struct {
	Apps []App `toml:"apps"`
}

type App struct {
	Name string `toml:"name"`
	Cmd  string `toml:"cmd"`
}

// LoadConfig reads the TOML file or creates a default one if it's missing
func LoadConfig() []App {
	home, err := os.UserHomeDir()
	if err != nil {
		return DefaultApps()
	}

	configDir := filepath.Join(home, ".config", "spymux")
	configFile := filepath.Join(configDir, "config.toml")

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		_ = os.MkdirAll(configDir, 0755)
		defaultToml := []byte("[[apps]]\nname = \"TERMINAL\"\ncmd = \"kitty --directory\"\n")
		_ = os.WriteFile(configFile, defaultToml, 0644)
		return DefaultApps()
	}

	fileData, err := os.ReadFile(configFile)
	if err != nil {
		return DefaultApps()
	}

	var conf AppConfig
	err = toml.Unmarshal(fileData, &conf)
	if err != nil || len(conf.Apps) == 0 {
		return DefaultApps()
	}

	return conf.Apps
}

func DefaultApps() []App {
	return []App{
		{Name: "TERMINAL", Cmd: "kitty --directory"},
	}
}
