package main

import (
	"errors"
	"fmt"
	"strings"
)

func Build(shape string, parameters ...float64) (Shape, error) {
	//	check if the shape is supported
	//	check if parameters are valid
	shape = strings.ToLower(strings.TrimSpace(shape))
	switch shape {
	case "rectangle":
		if len(parameters) != 2 || parameters[0] <= 0 || parameters[1] <= 0 {
			return nil, errors.New("invalid parameters")
		}
		return &Rectangle{Lside: parameters[0], Sside: parameters[1]}, nil
	case "square":
		if len(parameters) != 1 || parameters[0] <= 0 {
			return nil, errors.New("invalid parameters")
		}
		return &Square{Side: parameters[0]}, nil
	case "circle":
		if len(parameters) != 1 || parameters[0] <= 0 {
			return nil, errors.New("invalid parameters")
		}
		return &Circle{Radius: parameters[0]}, nil
	case "triangle":
		if len(parameters) != 1 || parameters[0] <= 0 {
			return nil, errors.New("invalid parameters")
		}
		return &EqTriangle{Side: parameters[0]}, nil
	default:
		return nil, errors.New("unsupported shape")
	}
}

func main() {
	//	call function in factory
	shape, err := Build("square", 6)
	if err != nil {
		fmt.Println(err)
	} else {
		Efficiency(shape)
	}
	shape, err = Build("rectangle", 6, 7)
	if err != nil {
		fmt.Println(err)
	} else {
		Efficiency(shape)
	}
	shape, err = Build("circle", -6)
	if err != nil {
		fmt.Println(err)
	} else {
		Efficiency(shape)
	}
	shape, err = Build("triangle", 6)
	if err != nil {
		fmt.Println(err)
	} else {
		Efficiency(shape)
	}

}
