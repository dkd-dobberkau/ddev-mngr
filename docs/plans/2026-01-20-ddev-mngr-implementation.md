# DDEV Manager Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Build a Go CLI tool to manage DDEV projects with interactive TUI and command-line modes.

**Architecture:** Cobra CLI framework handles commands, bubbletea powers the interactive TUI. A ddev client package wraps shell calls to `ddev` CLI and parses JSON output.

**Tech Stack:** Go 1.21+, Cobra (CLI), Bubbletea (TUI), Lipgloss (styling)

---

### Task 1: Initialize Go Module

**Files:**
- Create: `go.mod`
- Create: `main.go`

**Step 1: Initialize Go module**

Run:
```bash
go mod init github.com/oliverguenther/ddev-mngr
```
Expected: `go.mod` created

**Step 2: Create minimal main.go**

```go
package main

import "fmt"

func main() {
	fmt.Println("ddev-mngr")
}
```

**Step 3: Verify it compiles**

Run: `go build -o ddev-mngr && ./ddev-mngr`
Expected: Outputs "ddev-mngr"

**Step 4: Commit**

```bash
git add go.mod main.go
git commit -m "feat: initialize Go module"
```

---

### Task 2: DDEV Client - List Projects

**Files:**
- Create: `internal/ddev/client.go`
- Create: `internal/ddev/client_test.go`

**Step 1: Write the failing test**

```go
package ddev

import (
	"testing"
)

func TestParseListOutput(t *testing.T) {
	jsonOutput := `{"raw":[{"name":"test-project","status":"running","shortroot":"~/test","httpurl":"https://test.ddev.site"}]}`

	projects, err := ParseListOutput([]byte(jsonOutput))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(projects) != 1 {
		t.Fatalf("expected 1 project, got %d", len(projects))
	}

	if projects[0].Name != "test-project" {
		t.Errorf("expected name 'test-project', got '%s'", projects[0].Name)
	}

	if projects[0].Status != "running" {
		t.Errorf("expected status 'running', got '%s'", projects[0].Status)
	}
}
```

**Step 2: Run test to verify it fails**

Run: `go test ./internal/ddev/... -v`
Expected: FAIL - package not found

**Step 3: Write minimal implementation**

```go
package ddev

import (
	"encoding/json"
	"os/exec"
)

type Project struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	ShortRoot string `json:"shortroot"`
	HTTPUrl   string `json:"httpurl"`
}

type listOutput struct {
	Raw []Project `json:"raw"`
}

func ParseListOutput(data []byte) ([]Project, error) {
	var output listOutput
	if err := json.Unmarshal(data, &output); err != nil {
		return nil, err
	}
	return output.Raw, nil
}

func List() ([]Project, error) {
	cmd := exec.Command("ddev", "list", "--json-output")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return ParseListOutput(output)
}
```

**Step 4: Run test to verify it passes**

Run: `go test ./internal/ddev/... -v`
Expected: PASS

**Step 5: Commit**

```bash
git add internal/ddev/
git commit -m "feat: add DDEV client with list functionality"
```

---

### Task 3: DDEV Client - Start/Stop

**Files:**
- Modify: `internal/ddev/client.go`
- Modify: `internal/ddev/client_test.go`

**Step 1: Write the failing test**

Add to `internal/ddev/client_test.go`:

```go
func TestStartCommand(t *testing.T) {
	cmd := StartCommand("test-project")

	if cmd.Path == "" {
		t.Error("command path should not be empty")
	}

	args := cmd.Args
	// Args[0] is the command itself
	if len(args) < 3 || args[1] != "start" || args[2] != "-s" {
		t.Errorf("unexpected args: %v", args)
	}
}

func TestStopCommand(t *testing.T) {
	cmd := StopCommand("test-project")

	if cmd.Path == "" {
		t.Error("command path should not be empty")
	}

	args := cmd.Args
	if len(args) < 2 || args[1] != "stop" {
		t.Errorf("unexpected args: %v", args)
	}
}
```

**Step 2: Run test to verify it fails**

Run: `go test ./internal/ddev/... -v`
Expected: FAIL - StartCommand/StopCommand undefined

**Step 3: Write minimal implementation**

Add to `internal/ddev/client.go`:

```go
func StartCommand(name string) *exec.Cmd {
	return exec.Command("ddev", "start", "-s", name)
}

func StopCommand(name string) *exec.Cmd {
	return exec.Command("ddev", "stop", name)
}

func Start(name string) error {
	cmd := StartCommand(name)
	return cmd.Run()
}

func Stop(name string) error {
	cmd := StopCommand(name)
	return cmd.Run()
}
```

**Step 4: Run test to verify it passes**

Run: `go test ./internal/ddev/... -v`
Expected: PASS

**Step 5: Commit**

```bash
git add internal/ddev/
git commit -m "feat: add start/stop commands to DDEV client"
```

---

### Task 4: Cobra CLI Setup

**Files:**
- Create: `cmd/root.go`
- Modify: `main.go`

**Step 1: Install Cobra**

Run: `go get github.com/spf13/cobra`
Expected: go.mod updated

**Step 2: Create root command**

```go
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ddev-mngr",
	Short: "Manage DDEV projects",
	Long:  "A CLI tool to manage DDEV projects with interactive and command-line modes.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
```

**Step 3: Update main.go**

```go
package main

import "github.com/oliverguenther/ddev-mngr/cmd"

func main() {
	cmd.Execute()
}
```

**Step 4: Verify it compiles**

Run: `go build -o ddev-mngr && ./ddev-mngr --help`
Expected: Shows help text

**Step 5: Commit**

```bash
git add go.mod go.sum cmd/root.go main.go
git commit -m "feat: add Cobra CLI framework"
```

---

### Task 5: List Command

**Files:**
- Create: `cmd/list.go`

**Step 1: Create list command**

```go
package cmd

import (
	"fmt"
	"os"

	"github.com/oliverguenther/ddev-mngr/internal/ddev"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all DDEV projects",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := ddev.List()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		for _, p := range projects {
			symbol := "○"
			if p.Status == "running" {
				symbol = "●"
			}
			fmt.Printf("%s %-25s %s\n", symbol, p.Name, p.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
```

**Step 2: Verify it works**

Run: `go build -o ddev-mngr && ./ddev-mngr list`
Expected: Lists DDEV projects with status symbols

**Step 3: Commit**

```bash
git add cmd/list.go
git commit -m "feat: add list command"
```

---

### Task 6: Start Command

**Files:**
- Create: `cmd/start.go`

**Step 1: Create start command**

```go
package cmd

import (
	"fmt"
	"os"

	"github.com/oliverguenther/ddev-mngr/internal/ddev"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start <project>",
	Short: "Start a DDEV project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		fmt.Printf("Starting %s... ", name)

		if err := ddev.Start(name); err != nil {
			fmt.Println("failed")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("done")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
```

**Step 2: Verify it compiles**

Run: `go build -o ddev-mngr && ./ddev-mngr start --help`
Expected: Shows start command help

**Step 3: Commit**

```bash
git add cmd/start.go
git commit -m "feat: add start command"
```

---

### Task 7: Stop Command

**Files:**
- Create: `cmd/stop.go`

**Step 1: Create stop command**

```go
package cmd

import (
	"fmt"
	"os"

	"github.com/oliverguenther/ddev-mngr/internal/ddev"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop <project>",
	Short: "Stop a DDEV project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		fmt.Printf("Stopping %s... ", name)

		if err := ddev.Stop(name); err != nil {
			fmt.Println("failed")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("done")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
```

**Step 2: Verify it compiles**

Run: `go build -o ddev-mngr && ./ddev-mngr stop --help`
Expected: Shows stop command help

**Step 3: Commit**

```bash
git add cmd/stop.go
git commit -m "feat: add stop command"
```

---

### Task 8: Status Command

**Files:**
- Create: `cmd/status.go`

**Step 1: Create status command**

```go
package cmd

import (
	"fmt"
	"os"

	"github.com/oliverguenther/ddev-mngr/internal/ddev"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status [project]",
	Short: "Show status of DDEV projects",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := ddev.List()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		if len(args) == 1 {
			name := args[0]
			for _, p := range projects {
				if p.Name == name {
					fmt.Printf("%s: %s", p.Name, p.Status)
					if p.HTTPUrl != "" && p.Status == "running" {
						fmt.Printf(" (%s)", p.HTTPUrl)
					}
					fmt.Println()
					return
				}
			}
			fmt.Fprintf(os.Stderr, "Project '%s' not found\n", name)
			os.Exit(2)
		}

		// Show all
		running := 0
		stopped := 0
		for _, p := range projects {
			if p.Status == "running" {
				running++
			} else {
				stopped++
			}
		}
		fmt.Printf("Total: %d projects (%d running, %d stopped)\n", len(projects), running, stopped)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
```

**Step 2: Verify it works**

Run: `go build -o ddev-mngr && ./ddev-mngr status`
Expected: Shows project count summary

**Step 3: Commit**

```bash
git add cmd/status.go
git commit -m "feat: add status command"
```

---

### Task 9: TUI - Basic Model

**Files:**
- Create: `internal/tui/model.go`

**Step 1: Install bubbletea and lipgloss**

Run:
```bash
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
```
Expected: go.mod updated

**Step 2: Create TUI model**

```go
package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/oliverguenther/ddev-mngr/internal/ddev"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205"))

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("212"))

	runningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("42"))

	stoppedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)

type Model struct {
	projects []ddev.Project
	cursor   int
	loading  bool
	working  bool
	spinner  spinner.Model
	err      error
}

func NewModel() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return Model{
		spinner: s,
		loading: true,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, loadProjects)
}
```

**Step 3: Commit**

```bash
git add go.mod go.sum internal/tui/model.go
git commit -m "feat: add TUI model with styles"
```

---

### Task 10: TUI - Update Logic

**Files:**
- Create: `internal/tui/update.go`

**Step 1: Create update handler**

```go
package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/oliverguenther/ddev-mngr/internal/ddev"
)

type projectsLoadedMsg struct {
	projects []ddev.Project
	err      error
}

type actionDoneMsg struct {
	err error
}

func loadProjects() tea.Msg {
	projects, err := ddev.List()
	return projectsLoadedMsg{projects: projects, err: err}
}

func toggleProject(p ddev.Project) tea.Cmd {
	return func() tea.Msg {
		var err error
		if p.Status == "running" {
			err = ddev.Stop(p.Name)
		} else {
			err = ddev.Start(p.Name)
		}
		return actionDoneMsg{err: err}
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.working {
			return m, nil
		}

		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.projects)-1 {
				m.cursor++
			}
		case "enter":
			if len(m.projects) > 0 {
				m.working = true
				return m, tea.Batch(m.spinner.Tick, toggleProject(m.projects[m.cursor]))
			}
		case "r":
			m.loading = true
			return m, tea.Batch(m.spinner.Tick, loadProjects)
		}

	case projectsLoadedMsg:
		m.loading = false
		m.projects = msg.projects
		m.err = msg.err
		if m.cursor >= len(m.projects) {
			m.cursor = len(m.projects) - 1
		}
		if m.cursor < 0 {
			m.cursor = 0
		}

	case actionDoneMsg:
		m.working = false
		m.err = msg.err
		m.loading = true
		return m, loadProjects

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}
```

**Step 2: Commit**

```bash
git add internal/tui/update.go
git commit -m "feat: add TUI update logic"
```

---

### Task 11: TUI - View Rendering

**Files:**
- Create: `internal/tui/view.go`

**Step 1: Create view renderer**

```go
package tui

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("DDEV Project Manager"))
	b.WriteString("\n\n")

	if m.err != nil {
		b.WriteString(fmt.Sprintf("Error: %v\n\n", m.err))
	}

	if m.loading && len(m.projects) == 0 {
		b.WriteString(m.spinner.View() + " Loading projects...\n")
		return b.String()
	}

	if len(m.projects) == 0 {
		b.WriteString("No DDEV projects found.\n")
		return b.String()
	}

	for i, p := range m.projects {
		cursor := "  "
		if i == m.cursor {
			cursor = "> "
		}

		symbol := "○"
		style := stoppedStyle
		if p.Status == "running" {
			symbol = "●"
			style = runningStyle
		}

		line := fmt.Sprintf("%s%s %-25s %s", cursor, symbol, p.Name, p.Status)

		if i == m.cursor {
			if m.working {
				action := "Starting"
				if p.Status == "running" {
					action = "Stopping"
				}
				line = fmt.Sprintf("%s%s %-25s %s %s", cursor, symbol, p.Name, m.spinner.View(), action+"...")
			}
			b.WriteString(selectedStyle.Render(line))
		} else {
			b.WriteString(style.Render(line))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(helpStyle.Render("[↑↓/jk] Navigate  [Enter] Start/Stop  [r] Refresh  [q] Quit"))
	b.WriteString("\n")

	return b.String()
}
```

**Step 2: Commit**

```bash
git add internal/tui/view.go
git commit -m "feat: add TUI view rendering"
```

---

### Task 12: Integrate TUI into Root Command

**Files:**
- Modify: `cmd/root.go`

**Step 1: Update root command to launch TUI**

```go
package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/oliverguenther/ddev-mngr/internal/tui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ddev-mngr",
	Short: "Manage DDEV projects",
	Long:  "A CLI tool to manage DDEV projects with interactive and command-line modes.",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(tui.NewModel())
		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
```

**Step 2: Verify TUI works**

Run: `go build -o ddev-mngr && ./ddev-mngr`
Expected: Interactive TUI launches with project list

**Step 3: Commit**

```bash
git add cmd/root.go
git commit -m "feat: integrate TUI into root command"
```

---

### Task 13: Add .gitignore and README

**Files:**
- Create: `.gitignore`
- Create: `README.md`

**Step 1: Create .gitignore**

```
ddev-mngr
*.exe
.DS_Store
```

**Step 2: Create README**

```markdown
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
| ↑/k | Move up |
| ↓/j | Move down |
| Enter | Toggle start/stop |
| r | Refresh list |
| q/Esc | Quit |
```

**Step 3: Commit**

```bash
git add .gitignore README.md
git commit -m "docs: add README and .gitignore"
```

---

## Summary

13 tasks total:
1. Initialize Go module
2. DDEV client - list
3. DDEV client - start/stop
4. Cobra CLI setup
5. List command
6. Start command
7. Stop command
8. Status command
9. TUI - basic model
10. TUI - update logic
11. TUI - view rendering
12. Integrate TUI
13. Documentation
