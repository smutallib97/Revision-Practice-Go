package main

import "fmt"

func display[T any](x T) {
	fmt.Println(x)
}

func main() {
	display[string]("Focus on study")
	display[int](45)
	display[[]int]([]int{1, 2, 3, 4, 5})

}