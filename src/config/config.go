package config

import (
	"os"
	"spymux/src/utils"

	"github.com/pelletier/go-toml/v2"
)

type App struct {
	Name string `toml:"name"`
	Cmd  string `toml:"cmd"`
}

type Config struct {
	Spydir SpydirConfig `toml:"spydir"`
	Spybin SpybinConfig `toml:"spybin"`
}

type SpydirConfig struct {
	Apps []App `toml:"apps"`
}

type SpybinConfig struct {
	Apps []App `toml:"apps"`
}

func LoadConfig() (Config, error) {
	path, err := utils.ConfigFilePath("config.toml")
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var conf Config
	if err := toml.Unmarshal(data, &conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}
