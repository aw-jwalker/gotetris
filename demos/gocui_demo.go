package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// Left panel (stats)
	if v, err := g.SetView("stats", 0, 0, 20, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " Stats "
		v.Frame = true
		fmt.Fprintln(v, "Score: 0")
		fmt.Fprintln(v, "Level: 1")
		fmt.Fprintln(v, "Lines: 0")
	}

	// Center panel (game board)
	boardX := maxX/2 - 11
	if v, err := g.SetView("board", boardX, 0, boardX+22, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " Tetris "
		v.Frame = true
		fmt.Fprintln(v, "Game board")
		fmt.Fprintln(v, "would go here")
	}

	// Right panel (next pieces)
	if v, err := g.SetView("next", maxX-20, 0, maxX-1, 15); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " Next "
		v.Frame = true
		fmt.Fprintln(v, "Next pieces:")
	}

	// Bottom panel (controls)
	if v, err := g.SetView("controls", 0, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		fmt.Fprintln(v, "W=Drop | A/D=Move | Arrows=Rotate | Ctrl+C=Quit")
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
