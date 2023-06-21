package main

import (
	"strconv"
	"time"

	"github.com/gdamore/tcell"
)

const appleValue = 5

type Game struct {
	Screen    tcell.Screen
	snakeBody SneakBody
	apple     Apple
	score     int
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

func drawText(s tcell.Screen, text string, textStyle tcell.Style) {
	runes := []rune(text)
	s.SetContent(3, 3, ' ', runes, textStyle)
}

func drawScore(s tcell.Screen, scr int, scoreStyle tcell.Style) {
	text := strconv.Itoa(scr)
	text1 := "Score : " + text
	runes := []rune(text1)
	s.SetContent(4, 4, ' ', runes, scoreStyle)
}
func (g *Game) checkCollision(s tcell.Screen) {
	for _, part := range g.snakeBody.Parts {
		if g.apple.P == part.X && g.apple.T == part.Y {
			position := g.apple.randomPosition(s)
			g.apple.P = position.P
			g.apple.T = position.T
			g.score += appleValue
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
	textStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorGreen)
	scoreStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlue)
	g.apple.P = 10
	g.apple.T = 10
	g.score = 0

	for {
		g.Screen.Clear()
		DrawBorder(g.Screen, borderStyle, width, height)
		g.checkCollision(g.Screen)
		drawScore(g.Screen, g.score, scoreStyle)
		g.snakeBody.Update(width, height)
		drawParts(g.Screen, g.snakeBody.Parts, snakeStyle)
		drawApple(g.Screen, styleApple, g.apple)
		time.Sleep(60 * time.Millisecond)
		drawText(g.Screen, "Snake Game", textStyle)
		g.Screen.Show()
	}

}
