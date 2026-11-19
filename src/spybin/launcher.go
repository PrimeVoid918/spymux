package spybin

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func newLogger() (*log.Logger, func(), error) {
	stateDir, err := os.UserConfigDir()
	if err != nil {
		return nil, nil, err
	}

	logDir := filepath.Join(stateDir, "spymux", "logs")

	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, nil, err
	}

	logFile, err := os.OpenFile(
		filepath.Join(logDir, "app.log"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return nil, nil, err
	}

	logger := log.New(logFile, "[spymux] ", log.LstdFlags|log.Lshortfile)

	return logger, func() {
		_ = logFile.Close()
	}, nil
}

func launchApp(execCmd string) {
	logger, closeLog, err := newLogger()
	if err != nil {
		log.Printf("logger initialization failed: %v", err)
		return
	}
	defer closeLog()

	logger.Printf("Launching: %s", execCmd)

	if os.Getenv("HYPRLAND_INSTANCE_SIGNATURE") != "" {
		luaExpr := fmt.Sprintf(
			"hl.dsp.exec_cmd(%q)",
			execCmd,
		)

		output, err := exec.Command(
			"hyprctl",
			"dispatch",
			luaExpr,
		).CombinedOutput()

		if err != nil {
			logger.Printf(
				"Hyprland launch failed: %v | output: %s",
				err,
				output,
			)
			return
		}

		logger.Printf("Hyprland dispatch successful: %s", output)
		return
	}

	cmd := exec.Command("sh", "-c", execCmd)

	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Stdin = nil

	// Detach from the current terminal session.
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}

	if err := cmd.Start(); err != nil {
		logger.Printf("Direct launch failed: %v", err)
		return
	}

	logger.Printf("Process started: PID %d", cmd.Process.Pid)

	_ = cmd.Process.Release()
}
