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
	ready        bool
	width        int
	height       int
	score        int
	level        int
	lines        int
	board        *Board
	currentPiece *Piece
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
				renderBoardWithPiece(m.board, m.currentPiece, scale),
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

// renderBoardWithPiece renders the board with the current piece overlaid
func renderBoardWithPiece(board *Board, piece *Piece, scale int) string {
	if piece == nil {
		return board.Render(scale)
	}

	// Get the cells occupied by the current piece
	pieceCells := piece.Cells()
	pieceColor := piece.Color()

	// Create a temporary board copy to avoid mutating the original
	tempBoard := &Board{}
	for row := 0; row < BoardHeight; row++ {
		for col := 0; col < BoardWidth; col++ {
			tempBoard.Cells[row][col] = board.Cells[row][col]
		}
	}

	// Overlay the piece cells
	for _, cell := range pieceCells {
		if cell.Row >= 0 && cell.Row < BoardHeight &&
			cell.Col >= 0 && cell.Col < BoardWidth {
			tempBoard.Cells[cell.Row][cell.Col] = NewFilledCell(pieceColor)
		}
	}

	return tempBoard.Render(scale)
}

func main() {
	board := NewBoard()

	// Create a test piece (T piece near top-center for visibility)
	testPiece := NewPiece(PieceT, 2, 3)

	// Create the program with alt screen mode (fullscreen)
	p := tea.NewProgram(
		model{
			score:        0,
			level:        1,
			lines:        0,
			board:        board,
			currentPiece: testPiece,
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
