package main

import "fmt"

type Circle struct {
	r float64
}

func (c Circle) area() float64 {
	return 3.14 * c.r * c.r
}

func main() {

	circle1 := Circle{r: 2}

	fmt.Println(circle1.area())

}
