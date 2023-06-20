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
	apple := *a
	width, height := s.Size()
	maxp, maxt := width-3, height-3
	apple.P = rand.Intn(maxp + 1)
	apple.T = rand.Intn(maxt + 1)

	return apple
}
