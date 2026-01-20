---
layout: default
title: ddev-mngr
description: A CLI tool to manage DDEV projects with interactive TUI
---

# ddev-mngr

A fast, interactive CLI tool to manage your DDEV projects.

## Features

- **Interactive TUI** - Navigate with arrow keys, start/stop with Enter
- **CLI Mode** - Script-friendly commands for automation
- **Fast** - Built with Go, instant startup
- **Cross-platform** - macOS and Linux support

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

## Usage

### Interactive Mode

Simply run without arguments to launch the interactive TUI:

```bash
ddev-mngr
```

```
DDEV Project Manager

  ● my-project           running
  ○ another-project      stopped
> ● active-project       running

[↑↓] Navigate  [Enter] Start/Stop  [r] Refresh  [q] Quit
```

### CLI Mode

```bash
ddev-mngr list              # List all projects
ddev-mngr start <name>      # Start a project
ddev-mngr stop <name>       # Stop a project
ddev-mngr status [name]     # Show status
```

## Keybindings

| Key | Action |
|-----|--------|
| ↑/k | Move up |
| ↓/j | Move down |
| Enter | Toggle start/stop |
| r | Refresh list |
| q/Esc | Quit |

## Links

- [GitHub Repository](https://github.com/dkd-dobberkau/ddev-mngr)
- [Releases](https://github.com/dkd-dobberkau/ddev-mngr/releases)
- [Issues](https://github.com/dkd-dobberkau/ddev-mngr/issues)

## License

[Apache 2.0](https://github.com/dkd-dobberkau/ddev-mngr/blob/main/LICENSE)
