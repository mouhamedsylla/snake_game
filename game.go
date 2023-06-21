package main

import (
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	Screen    tcell.Screen
	snakeBody SneakBody
	apple     Apple
}

func drawParts(s tcell.Screen, parts []SnakePart, style tcell.Style) {
	for _, part := range parts {
		s.SetContent(part.X, part.Y, ' ', nil, style)
	}
}

func drawApple(s tcell.Screen, styleApple tcell.Style, a Apple) {
	x := a.P
	y := a.T
	s.SetContent(x, y, '🍎', nil, styleApple)
}
func (g *Game) checkCollision(s tcell.Screen) {
	for _, part := range g.snakeBody.Parts {
		if g.apple.P == part.X && g.apple.T == part.Y {
			position := g.apple.randomPosition(s)
			g.apple.P = position.P
			g.apple.T = position.T
		}
	}
}

func DrawBorder(s tcell.Screen, borderStyle tcell.Style, width int, height int) {
	s.SetContent(0, 0, '╔', nil, borderStyle)
	for i := 1; i <= width-2; i++ {
		s.SetContent(i, 0, '⠛', nil, borderStyle)
	}
	s.SetContent(width-1, 0, '╗', nil, borderStyle)
	s.SetContent(0, height-1, '╚', nil, borderStyle)
	for j := 1; j <= width-2; j++ {
		s.SetContent(j, height-1, '⠛', nil, borderStyle)
	}
	s.SetContent(width-1, height-1, '╝', nil, borderStyle)
	for k := 1; k <= height-2; k++ {
		s.SetContent(0, k, '▒', nil, borderStyle)
	}
	for l := 1; l <= height-2; l++ {
		s.SetContent(width-1, l, '▒', nil, borderStyle)
	}
}

func (g *Game) Run() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)
	g.Screen.SetStyle(defStyle)
	width, height := g.Screen.Size()
	snakeStyle := tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorGreen)
	borderStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)
	styleApple := tcell.StyleDefault

	g.apple.P = 10
	g.apple.T = 10

	for {
		g.Screen.Clear()
		g.snakeBody.Update(width, height)
		drawApple(g.Screen, styleApple, g.apple)
		drawParts(g.Screen, g.snakeBody.Parts, snakeStyle)
		DrawBorder(g.Screen, borderStyle, width, height)
		g.checkCollision(g.Screen)
		time.Sleep(40 * time.Millisecond)

		g.Screen.Show()
	}
}
