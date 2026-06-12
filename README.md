# Spymux

Spymux is a lightweight Linux launcher built with Go, Bubble Tea, and Lip Gloss. It combines two workflows into a single terminal application:

- **Spybin** — Application launcher powered by `.desktop` discovery.
- **Spydir** — Directory-aware workspace launcher powered by Zoxide.

Originally created as a personal replacement for a custom Rofi workflow, Spymux evolved into a dedicated launcher focused on keyboard-driven navigation, workspace spawning, and developer productivity.

## Features

### Spybin

- Scans system and user `.desktop` files
- Fuzzy application search
- Keyboard-driven navigation
- Automatic icon mapping for common applications
- Launch filtering for unwanted system entries
- Hyprland-aware process spawning

### Spydir

- Directory search powered by Zoxide rankings
- Horizontal application selection
- Vertical directory navigation
- Configurable application targets
- Workspace-oriented launch workflow
- Hyprland-aware process spawning

## Motivation

Spymux started as a collection of shell scripts built around FZF.

The original workflow worked well for launching applications, but eventually expanded into a second requirement:

- Select a frequently used directory
- Select a preferred tool
- Launch directly into a working environment

While FZF handled directory selection effectively, the interface was limited to a single navigation axis. Spydir was created to support simultaneous application and directory selection through a dedicated terminal user interface.

The project also became an opportunity to learn Go while building a tool used daily in a personal Linux environment.

## Modes

### Launcher Selection

Running Spymux without arguments opens the launcher selector.

    spymux

Available modes:

- SPYDIR (Directories)
- SPYBIN (Applications)

### Launch Spybin Directly

    spymux -b

### Launch Spydir Directly

    spymux -d

## Configuration

Configuration file:

    ~/.config/spymux/config.toml

Example:

    [[apps]]
    name = "TERMINAL"
    cmd = "kitty --directory {dir}"

    [[apps]]
    name = "NEOVIM"
    cmd = "kitty -d {dir} nvim"

    [[apps]]
    name = "CODE"
    cmd = "code {dir}"

The `{dir}` placeholder is replaced with the selected directory before execution.

## Recommended Hyprland Integration

For Hyprland users:

    bind = $mainMod SHIFT, return, exec, kitty --class spymux -e ~/.local/bin/spymux -d
    bind = $mainMod, space, exec, kitty --class spymux -e ~/.local/bin/spymux -b

Launching Spymux inside a dedicated terminal window provides a cleaner workflow and avoids leaving additional terminal processes open after spawning applications or workspaces.

## Dependencies

### Runtime

- Linux
- Zoxide (for Spydir)
- A terminal emulator such as Kitty

### Build

- Go
- Bubble Tea
- Lip Gloss

## Building

Build manually:

    go build -o spymux

Or use the included build script:

    ./build.sh

## Architecture

    spymux
    ├── spybin
    │   ├── Desktop entry discovery
    │   ├── Fuzzy search
    │   └── Application launching
    │
    ├── spydir
    │   ├── Zoxide integration
    │   ├── Directory ranking
    │   ├── Workspace spawning
    │   └── Configurable targets
    │
    └── tui
        └── Startup mode selection

## Technology Stack

- Go
- Bubble Tea
- Lip Gloss
- TOML Configuration
- Zoxide

## Future Improvements

- Spybin application launch profiles
- Custom launch arguments
- User-defined icon mappings
- Additional launcher layouts
- Improved filtering and ranking

## About

Spymux is a personal productivity tool developed to streamline application launching and workspace management on Linux. It serves as both a daily-driver utility and an exploration of Go-based terminal application development.
