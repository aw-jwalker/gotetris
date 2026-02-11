# TUI Library Demos

Compare gocui vs bubbletea side-by-side!

## gocui Demo

**Run:**

```bash
cd demos
go run gocui_demo.go
```

**Features:**

- Manual layout with absolute positioning
- Views created in layout function
- Lower-level control
- Lazygit-style

**Controls:** Ctrl+C to quit

## Bubbletea Demo

**Run:**

```bash
cd demos
go run bubbletea_demo.go
```

**Features:**

- Elm Architecture (model-update-view)
- Lipgloss styling (Kanagawa colors!)
- Higher-level abstraction
- Interactive demo

**Controls:**

- `Space` - Increase score
- `Up` - Increase level
- `Down` - Increase lines
- `Q` - Quit

## Comparison

**gocui:**

- ✅ What lazygit uses
- ✅ Direct control
- ❌ More code for styling
- ❌ Manual positioning

**bubbletea:**

- ✅ Cleaner code
- ✅ Better styling (lipgloss)
- ✅ Easier state management
- ✅ More examples online
- ✅ Already has Kanagawa colors in demo!

## Recommendation

**Use Bubbletea** - easier to work with, better ecosystem, great documentation.
