# GoTetris

A modern terminal Tetris game written in Go with a beautiful TUI interface
inspired by lazygit and sqlit.

## Features (Planned)

- 🎮 Classic Tetris gameplay
- 📊 Panel-based UI with:
  - Centered game board
  - Stats panel (score, level, lines)
  - Next pieces queue (3-5 pieces preview)
  - Controls help panel
  - High scores
- 🎨 Kanagawa color theme
- ⚡ Responsive terminal sizing
- 🎯 Smooth gameplay

## Tech Stack

- **Language**: Go
- **TUI Library**: Bubbletea (with lipgloss for styling)
- **Architecture**: Elm-inspired (model-update-view) with panel-based layout
- **Colors**: Kanagawa theme

## Development

```bash
# Run the game
go run main.go

# Build
go build -o gotetris

# Install
go install
```

## Project Status

🚧 **Work in Progress** - Building incrementally for learning and fun!

## Controls (Planned)

- `W` / `Space` - Hard drop
- `A` / `D` - Move left/right
- `S` - Soft drop
- `←` / `→` - Rotate
- `P` - Pause
- `Q` - Quit
