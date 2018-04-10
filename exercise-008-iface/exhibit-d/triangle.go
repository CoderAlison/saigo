package main

import (
	"math"
)

type EqTriangle struct {
	Side float64
}

func (t *EqTriangle) Name() string {
	return "Equilateral triangle"
}

func (t *EqTriangle) Perimeter() float64 {
	return 3 * t.Side
}

func (t *EqTriangle) Area() float64 {
	return math.Sqrt(3) / 4 * t.Side * t.Side
}
