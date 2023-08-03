package main

import "fmt"

func apply(numbers []int, F func(int) int) []int {
	res := make([]int, len(numbers))

	for i, number := range numbers {
		res[i] = F(number)
	}
	return res
}

func main() {
	numbers := []int{1, 2, 3, 5, 8, 7}

	res1 := apply(numbers, func(n int) int {
		return n * 2
	})
	fmt.Println(res1)
}
