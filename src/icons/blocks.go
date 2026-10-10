package icons

// BlockIcons holds all Unicode block elements, shades, halves, and quadrants for TUI design.
type Block struct {
	Shades     ShadeBlocks
	Vertical   VerticalBlocks
	Horizontal HorizontalBlocks
	Halves     HalfBlocks
	Quadrants  QuadrantBlocks
}

func newBlocks() Block {
	return Block{
		Shades:     newShades(),
		Vertical:   newVerticalBlocks(),
		Horizontal: newHorizontalBlocks(),
		Halves:     newHalves(),
		Quadrants:  newQuadrants(),
	}
}

// ShadeBlocks defines gradient shade patterns from light to full.
type ShadeBlocks struct {
	Light  string // ░ (U+2591)
	Medium string // ▒ (U+2592)
	Dark   string // ▓ (U+2593)
	Full   string // █ (U+2588)
}

func newShades() ShadeBlocks {
	return ShadeBlocks{
		Light:  "░",
		Medium: "▒",
		Dark:   "▓",
		Full:   "█",
	}
}

// VerticalBlocks defines vertical progress columns (left-aligned 1/8 to 8/8).
type VerticalBlocks struct {
	Full  string // █ (8/8)
	Bar78 string // ▉ (7/8)
	Bar34 string // ▊ (6/8 - Three Quarters)
	Bar58 string // ▋ (5/8)
	Half  string // ▌ (4/8 - Left Half)
	Bar38 string // ▍ (3/8)
	Bar14 string // ▎ (2/8 - One Quarter)
	Bar18 string // ▏ (1/8 - One Eighth)
}

func newVerticalBlocks() VerticalBlocks {
	return VerticalBlocks{
		Full:  "█",
		Bar78: "▉",
		Bar34: "▊",
		Bar58: "▋",
		Half:  "▌",
		Bar38: "▍",
		Bar14: "▎",
		Bar18: "▏",
	}
}

// HorizontalBlocks defines horizontal progress rows (bottom-aligned 1/8 to 8/8).
type HorizontalBlocks struct {
	Full  string // █ (8/8)
	Bar78 string // ▇ (7/8)
	Bar34 string // ▆ (6/8)
	Bar58 string // ▅ (5/8)
	Half  string // ▄ (4/8 - Lower Half)
	Bar38 string // ▃ (3/8)
	Bar14 string // ▂ (2/8)
	Bar18 string // ▁ (1/8)
}

func newHorizontalBlocks() HorizontalBlocks {
	return HorizontalBlocks{
		Full:  "█",
		Bar78: "▇",
		Bar34: "▆",
		Bar58: "▅",
		Half:  "▄",
		Bar38: "▃",
		Bar14: "▂",
		Bar18: "▁",
	}
}

// HalfBlocks defines major split halves and side lines.
type HalfBlocks struct {
	Upper   string // ▀ (Upper Half)
	Lower   string // ▄ (Lower Half)
	Right   string // ▐ (Right Half)
	Upper18 string // ▔ (Upper One Eighth)
	Left18  string // ▏ (Left One Eighth)
	Right18 string // ▕ (Right One Eighth)
}

func newHalves() HalfBlocks {
	return HalfBlocks{
		Upper:   "▀",
		Lower:   "▄",
		Right:   "▐",
		Upper18: "▔",
		Left18:  "▏",
		Right18: "▕",
	}
}

// QuadrantBlocks defines mosaic quadrants and triangle patterns (useful for custom shapes).
type QuadrantBlocks struct {
	LowerLeft      string // ▖
	LowerRight     string // ▗
	UpperLeft      string // ▘
	UpperLeftLower string // ▙
	Diagonal       string // ▚ (Checkerboard / Diagonal split)
	ThreeCornerTop string // ▛
	ThreeCornerBot string // ▜
	UpperRight     string // ▝
	InverseDiag    string // ▞
	ThreeCornerRt  string // ▟
}

func newQuadrants() QuadrantBlocks {
	return QuadrantBlocks{
		LowerLeft:      "▖",
		LowerRight:     "▗",
		UpperLeft:      "▘",
		UpperLeftLower: "▙",
		Diagonal:       "▚",
		ThreeCornerTop: "▛",
		ThreeCornerBot: "▜",
		UpperRight:     "▝",
		InverseDiag:    "▞",
		ThreeCornerRt:  "▟",
	}
}
