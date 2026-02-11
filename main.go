package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Styles using Kanagawa color scheme
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

// Model holds our application state
type model struct {
	ready  bool
	width  int
	height int
	score  int
	level  int
	lines  int
}

// Init is called once at startup
func (m model) Init() tea.Cmd {
	return nil
}

// Update handles incoming events and updates the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		// Handle keyboard input
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		// Handle terminal resize
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
	}

	return m, nil
}

// View renders the UI
func (m model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	// Calculate responsive dimensions
	// Account for borders (2 per panel), padding (4 per panel), and spacing (4 total)
	overhead := (3 * 6) + 4 // 3 panels * (2 border + 4 padding) + 4 spacing
	availableWidth := m.width - overhead
	if availableWidth < 60 {
		availableWidth = 60
	}

	sideWidth := availableWidth / 5
	if sideWidth < 16 {
		sideWidth = 16
	}
	boardWidth := availableWidth - (2 * sideWidth)
	if boardWidth < 20 {
		boardWidth = 20
	}

	boardHeight := m.height - 10
	if boardHeight < 10 {
		boardHeight = 10
	}

	// Stats panel with responsive width
	stats := statsStyle.Copy().
		Width(sideWidth).
		Render(
			titleStyle.Render("Stats") + "\n\n" +
				fmt.Sprintf("Score: %d\n", m.score) +
				fmt.Sprintf("Level: %d\n", m.level) +
				fmt.Sprintf("Lines: %d", m.lines),
		)

	// Board panel with responsive dimensions
	board := boardStyle.Copy().
		Width(boardWidth).
		Height(boardHeight).
		Render(
			titleStyle.Render("Tetris") + "\n\n" +
				"Game board\n" +
				"would go here",
		)

	// Next pieces panel with responsive width
	next := nextStyle.Copy().
		Width(sideWidth).
		Render(
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
		"Arrow Keys=Move | Space=Drop | Q=Quit",
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
	// Create the program with alt screen mode (fullscreen)
	p := tea.NewProgram(
		model{score: 0, level: 1, lines: 0},
		tea.WithAltScreen(),       // Fullscreen mode
		tea.WithMouseCellMotion(), // Mouse support
	)

	// Run the program
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
