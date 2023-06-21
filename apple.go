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
	minp, mint := 6, 6
	maxp, maxt := width-6, height-6
	randomApple.P = rand.Intn(maxp-minp+1) + minp
	randomApple.T = rand.Intn(maxt-mint+1) + mint

	return randomApple
}
