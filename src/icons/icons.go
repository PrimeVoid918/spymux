package icons

type Icons struct {
	System     System
	Layout     Layouts
	Apps       Apps
	Delimiters Delimiters
	Blocks     Block
}

func New() Icons {
	return Icons{
		System:     newSystem(),
		Layout:     newLayouts(),
		Apps:       newApps(),
		Delimiters: newDelimiters(),
		Blocks:     newBlocks(),
	}
}
