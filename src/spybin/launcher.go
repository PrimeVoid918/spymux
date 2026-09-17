package spybin

import (
	"os"
	"os/exec"
)

func launchApp(execCmd string) {
	if os.Getenv("HYPRLAND_INSTANCE_SIGNATURE") != "" {
		_ = exec.Command("hyprctl", "dispatch", "exec", "--", execCmd).Run()
		return
	}

	cmd := exec.Command("sh", "-c", execCmd)

	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Stdin = nil

	_ = cmd.Start() //! can return an error
	if cmd.Process != nil {
		_ = cmd.Process.Release()
	}
}
