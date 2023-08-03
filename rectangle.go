package main

import "fmt"

type Rectangle struct {
	length  float64
	breadth float64
}

//method with a value receiver
func (r Rectangle) displayAreaOfRectangle() {

	Area := r.length * r.breadth
	fmt.Println("Area of Rectangle", Area)

}

// method with a pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.breadth *= factor
	r.length *= factor
	fmt.Println(r.length, r.breadth)
}

func main() {

	rectangle1 := Rectangle{
		length:  4,
		breadth: 6,
	}

	rectangle2 := &Rectangle{
		length:  6,
		breadth: 4,
	}

	rectangle1.displayAreaOfRectangle()
	rectangle1.Scale(20)
	rectangle1.Scale(10)
	fmt.Println("Rectangle Value: ", rectangle1)
	rectangle2.Scale(20)
	rectangle2.Scale(10)

	fmt.Println("Rectangle 2 Value: ", rectangle2)
}
