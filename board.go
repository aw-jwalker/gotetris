package main

import "github.com/charmbracelet/lipgloss"

// CellColor represents the color of a cell on the board
type CellColor string

const (
	ColorEmpty  CellColor = ""
	ColorCyan   CellColor = "#7aa2f7" // I piece (Kanagawa blue)
	ColorYellow CellColor = "#e0af68" // O piece (Kanagawa yellow)
	ColorPurple CellColor = "#957fb8" // T piece (Kanagawa purple)
	ColorGreen  CellColor = "#76946a" // S piece (Kanagawa green)
	ColorRed    CellColor = "#e46876" // Z piece (Kanagawa red)
	ColorBlue   CellColor = "#7e9cd8" // J piece (Kanagawa primary blue)
	ColorOrange CellColor = "#ffa066" // L piece (Kanagawa orange)
)

// Cell represents a single cell on the board
type Cell struct {
	Filled bool
	Color  CellColor
}

// NewCell creates an empty cell
func NewCell() Cell {
	return Cell{Filled: false, Color: ColorEmpty}
}

// NewFilledCell creates a filled cell with a color
func NewFilledCell(color CellColor) Cell {
	return Cell{Filled: true, Color: color}
}

const (
	BoardWidth  = 10
	BoardHeight = 20
)

// Board represents the Tetris game board
type Board struct {
	Cells [BoardHeight][BoardWidth]Cell
}

// NewBoard creates a new empty board
func NewBoard() *Board {
	b := &Board{}
	for row := 0; row < BoardHeight; row++ {
		for col := 0; col < BoardWidth; col++ {
			b.Cells[row][col] = NewCell()
		}
	}
	return b
}

// SetCell sets a cell at the given position
func (b *Board) SetCell(row, col int, cell Cell) {
	if row >= 0 && row < BoardHeight && col >= 0 && col < BoardWidth {
		b.Cells[row][col] = cell
	}
}

// GetCell gets a cell at the given position
func (b *Board) GetCell(row, col int) Cell {
	if row >= 0 && row < BoardHeight && col >= 0 && col < BoardWidth {
		return b.Cells[row][col]
	}
	return NewCell()
}

// Render converts the board to a string for display with scaling
// scale determines how many terminal characters each cell uses
// scale=1: each cell is 2 chars wide × 1 line tall
// scale=2: each cell is 4 chars wide × 2 lines tall, etc.
func (b *Board) Render(scale int) string {
	if scale < 1 {
		scale = 1
	}

	var result string
	charsPerCell := 2 * scale // Each cell is 2 chars wide per scale unit

	for row := 0; row < BoardHeight; row++ {
		// Render each row 'scale' times vertically
		for lineInCell := 0; lineInCell < scale; lineInCell++ {
			for col := 0; col < BoardWidth; col++ {
				cell := b.Cells[row][col]
				if cell.Filled {
					// Render filled cell as colored block
					style := lipgloss.NewStyle().Foreground(lipgloss.Color(cell.Color))
					// Repeat the block character to fill the scaled width
					blocks := ""
					for i := 0; i < charsPerCell; i++ {
						blocks += "█"
					}
					result += style.Render(blocks)
				} else {
					// Render empty cell as dots
					style := lipgloss.NewStyle().Foreground(lipgloss.Color("#54546d")) // Dim gray from Kanagawa
					// Repeat the dot character to fill the scaled width
					dots := ""
					for i := 0; i < charsPerCell; i++ {
						dots += "·"
					}
					result += style.Render(dots)
				}
			}
			// Add newline after each line
			result += "\n"
		}
	}

	// Remove trailing newline
	if len(result) > 0 && result[len(result)-1] == '\n' {
		result = result[:len(result)-1]
	}

	return result
}
