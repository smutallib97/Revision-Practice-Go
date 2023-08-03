//Polymorphism concept is implemented in golang with the help of interface//

package main

import "fmt"

type Shape interface {
	Area() float64
}

type Square struct {
	Sides float64
}

type Circle struct {
	radius float64
}

func (s Square) Area() float64 {
	return 4 * s.Sides
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}
func main() {
	var squareShape Shape
	squareShape = Square{10}
	fmt.Println(squareShape.Area())

	var circleShape Shape

	circleShape = Circle{5}
	fmt.Println(circleShape.Area())
}
