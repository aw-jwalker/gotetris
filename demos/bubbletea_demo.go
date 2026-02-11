package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Styles
var (
	panelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7e9cd8"))

	statsStyle = panelStyle.Copy().
			Width(20).
			BorderForeground(lipgloss.Color("#76946a"))

	boardStyle = panelStyle.Copy().
			Width(24).
			Height(20).
			BorderForeground(lipgloss.Color("#7e9cd8"))

	nextStyle = panelStyle.Copy().
			Width(20).
			BorderForeground(lipgloss.Color("#957fb8"))

	controlsStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#c0a36e"))
)

type model struct {
	score int
	level int
	lines int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "space":
			m.score += 100
		case "up":
			m.level++
		case "down":
			m.lines++
		}
	}
	return m, nil
}

func (m model) View() string {
	// Stats panel
	stats := statsStyle.Render(
		titleStyle.Render("Stats") + "\n\n" +
			fmt.Sprintf("Score: %d\n", m.score) +
			fmt.Sprintf("Level: %d\n", m.level) +
			fmt.Sprintf("Lines: %d", m.lines),
	)

	// Board panel
	board := boardStyle.Render(
		titleStyle.Render("Tetris") + "\n\n" +
			"Game board\n" +
			"would go here",
	)

	// Next pieces panel
	next := nextStyle.Render(
		titleStyle.Render("Next") + "\n\n" +
			"Next pieces\n" +
			"preview",
	)

	// Layout panels horizontally
	top := lipgloss.JoinHorizontal(
		lipgloss.Top,
		stats,
		"  ",
		board,
		"  ",
		next,
	)

	// Controls at bottom
	controls := controlsStyle.Render(
		"Space=Drop | Up=Level | Down=Lines | Q=Quit",
	)

	// Join vertically
	return lipgloss.JoinVertical(
		lipgloss.Left,
		top,
		"\n",
		controls,
	)
}

func main() {
	p := tea.NewProgram(model{score: 0, level: 1, lines: 0})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
