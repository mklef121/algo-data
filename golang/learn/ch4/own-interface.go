package main

import (
	"fmt"
	"math"
)

type Shape2D interface {
	Perimeter() float64
}

type circle struct {
	R float64
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func main() {
	a := circle{R: 1.5}
	fmt.Printf("Radius %.2f -> Perimeter %.3f \n", a.R, a.Perimeter())

	// In this case, we examine whether a variable is of the Shape2D interface by using the interface{}(a). (Shape2D) notation
	_, ok := interface{}(a).(Shape2D)
	if ok {
		fmt.Println("a is a Shape2D!")
	}
}
