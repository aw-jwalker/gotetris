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
			PaddingTop(0).
			PaddingBottom(1).
			PaddingLeft(2).
			PaddingRight(2)

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
	board  *Board
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

	// Decide scale based on terminal size (1 or 2 only)
	scale := 1
	if m.width >= 60 && m.height >= 48 {
		scale = 2
	}

	// Calculate exact board dimensions at chosen scale
	// Base board: 10 cells wide (2 chars each) × 20 cells tall
	boardRenderWidth := 10 * 2 * scale // 20 or 40 chars
	boardRenderHeight := 20 * scale    // 20 or 40 lines

	// Panel dimensions: board + padding + title
	boardPanelWidth := boardRenderWidth + 4   // +4 for padding (2 on each side)
	boardPanelHeight := boardRenderHeight + 5 // +5 for title and padding

	// Side panel dimensions
	sideWidth := 20
	sideHeight := boardPanelHeight

	// Stats panel
	stats := statsStyle.Copy().
		Width(sideWidth).
		Height(sideHeight).
		AlignVertical(lipgloss.Top).
		Render(
			titleStyle.Render("Stats") + "\n\n" +
				fmt.Sprintf("Score: %d\n", m.score) +
				fmt.Sprintf("Level: %d\n", m.level) +
				fmt.Sprintf("Lines: %d", m.lines),
		)

	// Board panel sized exactly for the board
	board := boardStyle.Copy().
		Width(boardPanelWidth).
		Height(boardPanelHeight).
		AlignVertical(lipgloss.Top).
		Render(
			titleStyle.Render("Tetris") + "\n\n" +
				m.board.Render(scale),
		)

	// Next pieces panel
	next := nextStyle.Copy().
		Width(sideWidth).
		Height(sideHeight).
		AlignVertical(lipgloss.Top).
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

	// Join vertically (no extra spacing)
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		top,
		controls,
	)

	// Center horizontally, align to top
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Top,
		content,
	)
}

func main() {
	board := NewBoard()

	// Add some test filled cells for visual verification
	board.SetCell(19, 0, NewFilledCell(ColorRed)) // Bottom left
	board.SetCell(19, 1, NewFilledCell(ColorRed))
	board.SetCell(19, 2, NewFilledCell(ColorGreen))
	board.SetCell(18, 0, NewFilledCell(ColorBlue))
	board.SetCell(18, 1, NewFilledCell(ColorYellow))
	board.SetCell(17, 0, NewFilledCell(ColorPurple))
	board.SetCell(15, 5, NewFilledCell(ColorCyan))   // Middle
	board.SetCell(10, 9, NewFilledCell(ColorOrange)) // Top right area

	// Create the program with alt screen mode (fullscreen)
	p := tea.NewProgram(
		model{
			score: 0,
			level: 1,
			lines: 0,
			board: board,
		},
		tea.WithAltScreen(),       // Fullscreen mode
		tea.WithMouseCellMotion(), // Mouse support
	)

	// Run the program
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
