package main

import (
	"math/rand"
)

// Player ...
type RandoRex struct {
}

// Type returns the type of the player
func (r *RandoRex) Type() string {
	return "RandoRex"
}

// Play returns a move
func (r *RandoRex) Play() int {
	choice := rand.Int() % 3
	return choice
}
