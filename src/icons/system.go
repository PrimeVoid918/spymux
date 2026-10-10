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
	Git                string
	Code               string
	Script             string
	Package            string

	Status    StatusIcons
	Arrow     ArrowIcons
	Powerline PowerlineIcons
}

func newSystem() System {
	return System{
		Search:       "󰍉",
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
		Git:                "",
		Code:               "󰗀",
		Script:             "",
		Package:            "󰏗",

		Status:    newStatus(),
		Arrow:     newArrow(),
		Powerline: newPowerline(),
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

type PowerlineIcons struct {
	LeftBubble  string
	RightBubble string

	TriangleRight string
	TriangleLeft  string

	ChevronRight string
	ChevronLeft  string

	SlantRightTop    string
	SlantRightBottom string
	SlantLeftTop     string
	SlantLeftBottom  string
}

func newPowerline() PowerlineIcons {
	return PowerlineIcons{
		LeftBubble:  "",
		RightBubble: "",

		TriangleRight: "",
		TriangleLeft:  "",

		ChevronRight: "",
		ChevronLeft:  "",

		SlantRightTop:    "",
		SlantRightBottom: "",
		SlantLeftTop:     "",
		SlantLeftBottom:  "",
	}
}
