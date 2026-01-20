# ddev-mngr

A CLI tool to manage DDEV projects with interactive and command-line modes.

## Installation

### Homebrew (macOS/Linux)

```bash
brew tap dkd-dobberkau/tap
brew install ddev-mngr
```

### Go Install

```bash
go install github.com/dkd-dobberkau/ddev-mngr@latest
```

### Build from Source

```bash
git clone https://github.com/dkd-dobberkau/ddev-mngr.git
cd ddev-mngr
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
| ↑/k | Move up |
| ↓/j | Move down |
| Enter | Toggle start/stop |
| r | Refresh list |
| q/Esc | Quit |

## License

Apache 2.0
