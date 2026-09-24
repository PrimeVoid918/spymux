package icons

type System struct {
	Search       string
	ChevronRight string
	Divider      string
	ThickDivider string

	Target string
	Add    string
	Remove string

	Bluetooth          string
	File               string
	FileDocument       string
	DirectoryLocked    string
	DirectoryClose     string
	DirectoryOpen      string
	DirectoryEmpty     string
	DirectoryEmptyOpen string
	Executable         string
	Terminal           string
	Configuration      string
	Code               string
	Package            string

	Status StatusIcons
	Arrow  ArrowIcons
}

func newSystem() System {
	return System{
		Search:       "󰍉",
		ChevronRight: "󰅂",
		Divider:      "─",
		ThickDivider: "━",

		Target: "󰓾",
		Add:    "󰐕",
		Remove: "󰍴",

		Bluetooth:          "󰂯",
		File:               "󰈔",
		FileDocument:       "󰈙",
		DirectoryLocked:    "󰉐",
		DirectoryClose:     "󰉋",
		DirectoryOpen:      "󰝰",
		DirectoryEmpty:     "󰉖",
		DirectoryEmptyOpen: "󰷏 ",
		Executable:         "󰞷",
		Terminal:           "󰆍",
		Configuration:      "󰒓",
		Code:               "󰗀",
		Package:            "󰏗",

		Status: newStatus(),
		Arrow:  newArrow(),
	}
}

type StatusIcons struct {
	Check                string
	Cross                string
	CautionCircle        string
	CautionCircleSolid   string
	CautionTriangleSolid string
	CautionTriangle      string
	QuestionMark         string
	QuestionMarkSolid    string
}

func newStatus() StatusIcons {
	return StatusIcons{
		Check:                "󰄬",
		Cross:                "󰅖",
		CautionCircle:        "󰋽",
		CautionCircleSolid:   "󰀨",
		CautionTriangleSolid: "󰀪",
		CautionTriangle:      "󰀦",
		QuestionMark:         "󰋖",
		QuestionMarkSolid:    "󰋗",
	}
}

type ArrowIcons struct {
	Up           string
	Down         string
	Left         string
	Right        string
	DoubleUp     string
	DoubleDown   string
	DoubleLeft   string
	DoubleRight  string
	SolidUp      string
	SolidDown    string
	SolidLeft    string
	SolidRight   string
	OutlineUp    string
	OutlineDown  string
	OutlineLeft  string
	OutlineRight string
	UpLeft       string
	UpRight      string
	DownRight    string
	DownLeft     string
}

func newArrow() ArrowIcons {
	return ArrowIcons{
		Up:           "󰁝",
		Down:         "󰁅",
		Left:         "󰁍",
		Right:        "󰁔",
		DoubleUp:     "⇑",
		DoubleDown:   "⇓",
		DoubleLeft:   "⇐",
		DoubleRight:  "⇒",
		SolidUp:      "▲",
		SolidDown:    "▼",
		SolidLeft:    "◀",
		SolidRight:   "▶",
		OutlineUp:    "△ ",
		OutlineDown:  "▽",
		OutlineLeft:  "◁",
		OutlineRight: "▷",
		UpLeft:       "↖",
		UpRight:      "↗",
		DownRight:    "↘",
		DownLeft:     "↙",
	}
}
