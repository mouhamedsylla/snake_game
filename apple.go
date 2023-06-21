package main

import (
	"math/rand"

	"github.com/gdamore/tcell"
)

type Apple struct {
	P int
	T int
}

func (a *Apple) randomPosition(s tcell.Screen) Apple {
	randomApple := *a
	width, height := s.Size()
	maxp, maxt := width-3, height-3
	randomApple.P = rand.Intn(maxp + 1)
	randomApple.T = rand.Intn(maxt + 1)

	return randomApple
}
