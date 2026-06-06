func main() {
	p := tea.NewProgram(initialModel())
	p.Run()
}

func launch(mode, dir string) {
	switch mode {

	case "TERMINAL":
		exec.Command("hyprctl", "dispatch", "exec",
			"kitty --directory "+dir).Run()

	case "TMUX":
		exec.Command("hyprctl", "dispatch", "exec",
			"kitty -d "+dir+" tmux").Run()

	case "CODE-OSS":
		exec.Command("hyprctl", "dispatch", "exec",
			"code-oss "+dir).Run()

	case "NEOVIM":
		exec.Command("hyprctl", "dispatch", "exec",
			"kitty -d "+dir+" nvim").Run()
	}
}

