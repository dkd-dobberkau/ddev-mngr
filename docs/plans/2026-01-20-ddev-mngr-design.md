# DDEV Manager (ddev-mngr) Design

## Overview

A Go CLI tool for managing DDEV projects with both command-line and interactive modes.

## Features

- **Start/Stop/Status** - Core project management
- **CLI Mode** - Arguments for scripting: `ddev-mngr start <name>`
- **Interactive Mode** - Arrow-key navigation without arguments

## Architecture

### Commands

```
ddev-mngr                    # Interactive mode
ddev-mngr list               # List all projects
ddev-mngr start <name>       # Start project
ddev-mngr stop <name>        # Stop project
ddev-mngr status [name]      # Show status
```

### Project Structure

```
ddev-mngr/
├── main.go              # Entry point, CLI parsing
├── cmd/
│   ├── root.go          # Cobra root command
│   ├── list.go          # list subcommand
│   ├── start.go         # start subcommand
│   ├── stop.go          # stop subcommand
│   └── status.go        # status subcommand
├── internal/
│   ├── ddev/
│   │   └── client.go    # DDEV CLI wrapper
│   └── tui/
│       ├── model.go     # Bubbletea model
│       ├── view.go      # Rendering
│       └── update.go    # Event handling
├── go.mod
└── go.sum
```

### Dependencies

- `github.com/spf13/cobra` - CLI framework
- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/lipgloss` - Terminal styling

### DDEV Integration

- Data source: `ddev list --json-output`
- Start: `ddev start -s <name>` (from any directory)
- Stop: `ddev stop <name>`

## Interactive Mode (TUI)

### Display

```
DDEV Project Manager

  ● ai-usage              running
  ○ conceptdetection-demo stopped
  ● ddev-python           running
> ○ drupal-cms-1          stopped    ← Cursor
  ● kreditkarten          running

[↑↓] Navigate  [Enter] Start/Stop  [q] Quit
```

### Keybindings

| Key | Action |
|-----|--------|
| ↑/k | Move up |
| ↓/j | Move down |
| Enter | Toggle start/stop |
| r | Refresh list |
| q/Esc | Quit |

### Behavior

- `●` green = running, `○` gray = stopped
- Enter on running project → stops it
- Enter on stopped project → starts it
- During action: spinner + "Starting..." / "Stopping..."
- After action: auto-refresh list

## CLI Mode

### Examples

```bash
# List all projects
$ ddev-mngr list
● ai-usage       running
○ drupal-cms-1   stopped

# Start project
$ ddev-mngr start drupal-cms-1
Starting drupal-cms-1... done

# Stop project
$ ddev-mngr stop ai-usage
Stopping ai-usage... done

# Status of project
$ ddev-mngr status ai-usage
ai-usage: running (https://ai-usage.ddev.site)
```

## Error Handling

| Scenario | Response |
|----------|----------|
| DDEV not installed | Error message + exit 1 |
| Project not found | "Project 'xyz' not found" |
| Start/Stop failed | Pass through DDEV error |

### Exit Codes

- 0 = Success
- 1 = General error
- 2 = Project not found
