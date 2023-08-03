package main

import "fmt"

func main() {
	// Declare a variable and assign a value
	var num int = 42

	// Declare a pointer variable that points to the memory address of 'num'
	var ptr *int = &num

	// Printing the value and memory address of 'num'
	fmt.Printf("Value of 'num': %d\n", num)
	fmt.Printf("Memory Address of 'num': %p\n", &num)

	// Printing the value and memory address stored in the pointer 'ptr'
	fmt.Printf("Value of 'ptr': %p\n", ptr)
	fmt.Printf("Memory Address stored in 'ptr': %p\n", &ptr)

	// Dereferencing the pointer to access the value it points to
	fmt.Printf("Dereferenced value of 'ptr': %d\n", *ptr)

	// Modifying the value of 'num' through the pointer 'ptr'
	*ptr = 100
	fmt.Printf("Modified value of 'num': %d\n", num)
}
