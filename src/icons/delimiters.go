package icons

type Delimiters struct {
	Brackets    BracketPair
	Parentheses ParenPair
	Braces      BracePair
	Angles      AnglePair
	Separators  SeparatorIcons
}

func newDelimiters() Delimiters {
	return Delimiters{
		Brackets:    BracketPair{Left: "[", Right: "]"},
		Parentheses: ParenPair{Left: "(", Right: ")"},
		Braces:      BracePair{Left: "{", Right: "}"},
		Angles:      AnglePair{Left: "<", Right: ">"},
		Separators:  SeparatorIcons{Pipe: "|", Slash: "/", Backslash: "\\"},
	}
}

type BracketPair struct{ Left, Right string }
type ParenPair struct{ Left, Right string }
type BracePair struct{ Left, Right string }
type AnglePair struct{ Left, Right string }

type SeparatorIcons struct {
	Pipe      string
	Slash     string
	Backslash string
}
