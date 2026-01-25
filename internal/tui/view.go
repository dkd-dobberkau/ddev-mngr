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

		urlInfo := ""
		if p.Status == "running" && p.HTTPUrl != "" {
			urlInfo = fmt.Sprintf("  %s", p.HTTPUrl)
		}

		line := fmt.Sprintf("%s%s %-25s %-10s%s", cursor, symbol, p.Name, p.Status, urlInfo)

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

	if m.poweringOff {
		b.WriteString(m.spinner.View() + " Stopping all projects...\n")
	} else if m.confirmingPoweroff {
		b.WriteString(selectedStyle.Render("Press p again to stop all projects, any other key to cancel"))
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(helpStyle.Render("[↑↓/jk] Navigate  [Enter] Start/Stop  [p] Poweroff  [r] Refresh  [q] Quit"))
	b.WriteString("\n")

	return b.String()
}
