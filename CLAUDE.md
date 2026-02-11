# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with
code in this repository.

## Project Overview

GoTetris is a terminal-based Tetris game built with Go using the Bubbletea TUI
framework. The project follows an **Elm Architecture** pattern
(model-update-view) inspired by lazygit and sqlit, with a panel-based UI layout.

## Commands

```bash
# Run the game
go run main.go

# Build the executable
go build -o gotetris

# Install to GOPATH
go install

# Run demo comparisons (from demos/ directory)
cd demos
go run bubbletea_demo.go  # Recommended TUI approach
go run gocui_demo.go      # Alternative TUI library
```

## Git Workflow

### Branch Naming Convention

Always create a GitHub issue for planned work, then branch off `main` with this
naming pattern:

```
gh-{issue#}-{short-description}
```

- Prefix with `gh-` followed by the issue number
- Keep description short (ideally two words)
- Use hyphens between words
- Examples: `gh-1-add-panels`, `gh-5-fix-rotation`, `gh-12-score-system`

This convention makes it easy to review diffs against main and track work back
to issues.

### Commit Strategy

Make **small, incremental commits** throughout development:

- Each commit should represent a single logical change
- Commit frequently as you make progress
- This makes it easy to scroll through commit history and understand changes
  over time
- Helps with debugging by isolating when specific changes were introduced

**Use Conventional Commits format:**

```
<type>: <description>
```

Common types:

- `feat:` - New feature (e.g., `feat: add panel layout structure`)
- `fix:` - Bug fix (e.g., `fix: correct piece rotation logic`)
- `refactor:` - Code restructuring without behavior change (e.g.,
  `refactor: extract styling utilities`)
- `chore:` - Maintenance tasks (e.g., `chore: update dependencies`)
- `docs:` - Documentation changes (e.g., `docs: update README with controls`)
- `style:` - Code formatting/styling (e.g., `style: format with gofmt`)
- `test:` - Adding or updating tests (e.g., `test: add board collision tests`)
- `perf:` - Performance improvements (e.g., `perf: optimize render loop`)

## Architecture

### Bubbletea (Elm Architecture)

The application uses Bubbletea's Elm-inspired pattern with three core functions:

- **Init()**: One-time initialization, returns commands
- **Update(msg)**: Handles all events (keyboard, terminal resize, etc.), returns
  updated model and commands
- **View()**: Renders the current model state to a string

### Model Structure

The model holds all application state. Current minimal structure:

```go
type model struct {
    ready bool  // Whether terminal size is known
}
```

Future models will include game state (board, pieces, score, etc.)

### Panel-Based Layout

The UI uses **lipgloss** for styling and layout with a multi-panel design:

- Stats panel (score, level, lines)
- Centered game board
- Next pieces queue (3-5 preview)
- Controls help panel

Panels are composed using `lipgloss.JoinHorizontal()` and
`lipgloss.JoinVertical()`.

### Color Theme

Uses the **Kanagawa** color scheme. Key colors from the demo:

- `#7e9cd8` - Primary (blue)
- `#76946a` - Green
- `#957fb8` - Purple
- `#c0a36e` - Gold

Define styles as package-level variables with `lipgloss.NewStyle()`.

### Terminal Handling

- **Alt Screen Mode**: Use `tea.WithAltScreen()` for fullscreen mode (doesn't
  pollute terminal history)
- **Window Resize**: Handle `tea.WindowSizeMsg` in Update() to adjust layout
- **Mouse Support**: Optional via `tea.WithMouseCellMotion()`

## Project Structure

- `main.go` - Main application entry point with Bubbletea setup
- `demos/` - TUI library comparison demos (bubbletea vs gocui)
- Flat structure currently; no `/pkg` or `/internal` directories yet

## Design Decisions

- **Bubbletea over gocui**: Chose Bubbletea for cleaner code, better styling
  with lipgloss, easier state management, and richer ecosystem
- **Panel-based layout**: Modular UI components that can be independently styled
  and positioned
- **Elm Architecture**: Predictable state management with clear separation
  between model, updates, and view
