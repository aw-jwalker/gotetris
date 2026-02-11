package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Model holds our application state
type model struct {
	ready bool
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
		m.ready = true
	}

	return m, nil
}

// View renders the UI
func (m model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	return "GoTetris\n\nPress 'q' to quit"
}

func main() {
	// Create the program with alt screen mode (fullscreen)
	p := tea.NewProgram(
		model{},
		tea.WithAltScreen(),       // Fullscreen mode
		tea.WithMouseCellMotion(), // Mouse support
	)

	// Run the program
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
