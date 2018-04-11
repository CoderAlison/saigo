package main

type Obsessed struct {
	obsession int
}

func (o *Obsessed) Type() string {
	return "Obsessed"
}

func (o *Obsessed) Play() int {
	choice := o.obsession
	return choice
}
