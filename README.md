# ddev-mngr

A CLI tool to manage DDEV projects with interactive and command-line modes.

## Installation

```bash
go install github.com/oliverguenther/ddev-mngr@latest
```

Or build from source:

```bash
go build -o ddev-mngr
```

## Usage

### Interactive Mode

```bash
ddev-mngr
```

Use arrow keys to navigate, Enter to start/stop projects.

### CLI Mode

```bash
ddev-mngr list              # List all projects
ddev-mngr start <name>      # Start a project
ddev-mngr stop <name>       # Stop a project
ddev-mngr status [name]     # Show status
```

## Keybindings (Interactive Mode)

| Key | Action |
|-----|--------|
| Up/k | Move up |
| Down/j | Move down |
| Enter | Toggle start/stop |
| r | Refresh list |
| q/Esc | Quit |
