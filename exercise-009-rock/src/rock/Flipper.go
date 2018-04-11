package main

import (
	"math/rand"
)

type Flipper struct {
	choice1 int
	choice2 int
}

func (f *Flipper) Type() string {
	return "Flipper"
}

func (f *Flipper) Play() int {
	choice := rand.Int() % 2
	if choice == 0 {
		return f.choice1
	} else {
		return f.choice2
	}
}
