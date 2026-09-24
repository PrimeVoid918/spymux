package icons

type Apps struct {
	Browser     string
	Game        string
	Terminal    string
	CodeEditor  string
	Docs        string
	Settings    string
	Notes       string
	FileManager string
	Email       string
	DBMS        string
	Torrent     string
	Downloader  string
	ImageEditor string
	ImageViewer string
	VideoPlayer string
	VideoEditor string
	Camera      string
	Audio       string
	Capture     string
	Bluetooth   string
	Default     string
}

func newApps() Apps {
	return Apps{
		Browser:     "󰖟",
		Game:        "󰊴",
		Terminal:    "󰆍",
		CodeEditor:  "󰗀",
		Docs:        "󰈬",
		Settings:    "󰒓",
		Notes:       "󰠮",
		FileManager: "󰉋",
		Email:       "󰇮",
		DBMS:        "󰆼",
		Torrent:     "󱘖",
		Downloader:  "󰇚",
		ImageEditor: "󰏘",
		ImageViewer: "󰋩",
		VideoPlayer: "󰕼",
		VideoEditor: "󰚺",
		Camera:      "󰄀",
		Audio:       "󰕾",
		Capture:     "󰕧",
		Bluetooth:   "󰂯",
		Default:     "󰏗",
	}
}
