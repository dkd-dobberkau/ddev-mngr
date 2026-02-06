package tui

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dkd-dobberkau/ddev-mngr/internal/ddev"
)

type projectsLoadedMsg struct {
	projects []ddev.Project
	err      error
}

type actionDoneMsg struct {
	err error
}

type poweroffDoneMsg struct {
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

func poweroffAll() tea.Msg {
	err := ddev.Poweroff()
	return poweroffDoneMsg{err: err}
}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	return cmd.Start()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.working || m.poweringOff {
			return m, nil
		}

		if m.confirmingPoweroff {
			if msg.String() == "p" {
				m.confirmingPoweroff = false
				m.poweringOff = true
				return m, tea.Batch(m.spinner.Tick, poweroffAll)
			}
			m.confirmingPoweroff = false
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
		case "p":
			m.confirmingPoweroff = true
			return m, nil
		case "o":
			if len(m.projects) > 0 {
				p := m.projects[m.cursor]
				if p.Status == "running" && p.HTTPSUrl != "" {
					openBrowser(p.HTTPSUrl)
				}
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

	case poweroffDoneMsg:
		m.poweringOff = false
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
