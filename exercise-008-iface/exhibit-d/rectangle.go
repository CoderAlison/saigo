package main

type Rectangle struct {
	Lside float64
	Sside float64
}

func (r *Rectangle) Name() string {
	return "Rectangle"
}

func (r *Rectangle) Perimeter() float64 {
	return 2*r.Lside + 2*r.Sside
}

func (r *Rectangle) Area() float64 {
	return r.Lside * r.Sside
}
