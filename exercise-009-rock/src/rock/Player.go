package main

type Player interface {
	Type() string
	Play() int
}

func Play(p Player) int {
	return p.Play()
}
