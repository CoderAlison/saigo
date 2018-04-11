package main

// Player ...
type Cyclone struct {
	counter int
}

// Type returns the type of the player
func (c *Cyclone) Type() string {
	return "Cyclone"
}

// Play returns a move
func (c *Cyclone) Play() int {
	choice := c.counter % 3
	c.counter = (c.counter + 1) % 3
	return choice
}
