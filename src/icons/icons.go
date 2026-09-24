package icons

type Icons struct {
	System System
	Layout Layouts
	Apps   Apps
}

func New() Icons {
	return Icons{
		System: newSystem(),
		Layout: newLayouts(),
		Apps:   newApps(),
	}
}
