package config

import (
	"encoding/json"
	"os"
	utils "spymux/src/utils"

	"github.com/charmbracelet/lipgloss"
)

type walSpecial struct {
	Background string `json:"background"`
	Foreground string `json:"foreground"`
	Cursor     string `json:"cursor"`
	Accent     string `json:"accent"`
	SubAccent  string `json:"sub_accent"`
}

type walColors struct {
	Color0, Color1, Color2, Color3     string
	Color4, Color5, Color6, Color7     string
	Color8, Color9, Color10, Color11   string
	Color12, Color13, Color14, Color15 string
}

type walSchema struct {
	Special walSpecial `json:"special"`
	Colors  walColors  `json:"colors"`
}

type AppTheme struct {
	bg        lipgloss.Color
	fg        lipgloss.Color
	accent    lipgloss.Color
	subAccent lipgloss.Color
	palette   [16]lipgloss.Color
}

func defaultTheme() *AppTheme {
	return &AppTheme{
		bg:        lipgloss.Color("#000000"),
		fg:        lipgloss.Color("#B7B7B7"),
		accent:    lipgloss.Color("#B0B0B0"),
		subAccent: lipgloss.Color("#B0B0B0"),
		palette: [16]lipgloss.Color{
			lipgloss.Color("#010101"), lipgloss.Color("#434343"), lipgloss.Color("#646464"), lipgloss.Color("#747474"),
			lipgloss.Color("#7D7D7D"), lipgloss.Color("#B7B7B7"), lipgloss.Color("#B0B0B0"), lipgloss.Color("#B7B7B7"),
			lipgloss.Color("#808080"), lipgloss.Color("#434343"), lipgloss.Color("#646464"), lipgloss.Color("#747474"),
			lipgloss.Color("#7D7D7D"), lipgloss.Color("#B7B7B7"), lipgloss.Color("#B0B0B0"), lipgloss.Color("#B7B7B7"),
		},
	}
}

func LoadSystemTheme() (*AppTheme, error) {
	walCachePath, pathErr := utils.WalCacheDir()
	if pathErr != nil {
		return defaultTheme(), pathErr
	}

	fileData, err := os.ReadFile(walCachePath)
	if err != nil {
		if os.IsNotExist(err) {
			return defaultTheme(), nil
		}
		return defaultTheme(), err
	}

	var raw walSchema
	if err := json.Unmarshal(fileData, &raw); err != nil {
		return defaultTheme(), err
	}

	return &AppTheme{
		bg:        lipgloss.Color(raw.Special.Background),
		fg:        lipgloss.Color(raw.Special.Foreground),
		accent:    lipgloss.Color(raw.Colors.Color12),
		subAccent: lipgloss.Color(raw.Colors.Color7),
		palette: [16]lipgloss.Color{
			lipgloss.Color(raw.Colors.Color0), lipgloss.Color(raw.Colors.Color1), lipgloss.Color(raw.Colors.Color2), lipgloss.Color(raw.Colors.Color3),
			lipgloss.Color(raw.Colors.Color4), lipgloss.Color(raw.Colors.Color5), lipgloss.Color(raw.Colors.Color6), lipgloss.Color(raw.Colors.Color7),
			lipgloss.Color(raw.Colors.Color8), lipgloss.Color(raw.Colors.Color9), lipgloss.Color(raw.Colors.Color10), lipgloss.Color(raw.Colors.Color11),
			lipgloss.Color(raw.Colors.Color12), lipgloss.Color(raw.Colors.Color13), lipgloss.Color(raw.Colors.Color14), lipgloss.Color(raw.Colors.Color15),
		},
	}, nil
}

func (t *AppTheme) BG() lipgloss.Color        { return t.bg }
func (t *AppTheme) FG() lipgloss.Color        { return t.fg }
func (t *AppTheme) Accent() lipgloss.Color    { return t.accent }
func (t *AppTheme) SubAccent() lipgloss.Color { return t.subAccent }

func (t *AppTheme) Color(index int) lipgloss.Color {
	if index < 0 || index > 15 {
		return t.fg
	}
	return t.palette[index]
}
