package icons

type Layouts struct {
	Lines
}

func newLayouts() Layouts {
	return Layouts{
		Lines: newLines(),
	}
}

type Lines struct {
	TopLeft     string
	TopRight    string
	BottomRight string
	BottomLeft  string

	Horizontal       string
	Vertical         string
	SharpTopLeft     string
	SharpTopRight    string
	SharpBottomRight string
	SharpBottomLeft  string

	TopBottomLeftJunction   string
	TopBottomRightJunction  string
	LeftRightBottomJunction string
	LeftRightTopJunction    string
	CenterJunction          string

	ThickHorizontal       string
	ThickVertical         string
	ThickSharpTopLeft     string
	ThickSharpTopRight    string
	ThickSharpBottomRight string
	ThickSharpBottomLeft  string

	ThickTopBottomLeftJunction   string
	ThickTopBottomRightJunction  string
	ThickLeftRightBottomJunction string
	ThickLeftRightTopJunction    string
	ThickCenterJunction          string
}

func newLines() Lines {
	return Lines{
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomRight: "╯",
		BottomLeft:  "╰",

		Horizontal:       "─",
		Vertical:         "│",
		SharpTopLeft:     "┌",
		SharpTopRight:    "┐",
		SharpBottomRight: "┘",
		SharpBottomLeft:  "└",

		TopBottomLeftJunction:   "├",
		TopBottomRightJunction:  "┤",
		LeftRightBottomJunction: "┬ ",
		LeftRightTopJunction:    "┴ ┼",
		CenterJunction:          "┼",

		ThickHorizontal:       "┃",
		ThickVertical:         "━",
		ThickSharpTopLeft:     "┏",
		ThickSharpTopRight:    "┓",
		ThickSharpBottomRight: "┗",
		ThickSharpBottomLeft:  "┛",

		ThickTopBottomLeftJunction:   "┣",
		ThickTopBottomRightJunction:  "┫",
		ThickLeftRightBottomJunction: "┳",
		ThickLeftRightTopJunction:    "┻",
		ThickCenterJunction:          "╋",
	}
}
