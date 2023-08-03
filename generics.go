/*The interface{} Type:
In Go, the interface{} type serves as a generic type or a way to represent any data type.
It allows you to work with values of unknown type, which is useful when dealing with
heterogeneous collections or when you want to write more flexible code.
Since interface{} can hold any value, it can represent values of any type,
including built-in types, user-defined types, structs, pointers, and more.
Functions can take interface{} as a parameter, making them able to accept values of any type.
This flexibility comes at the cost of type safety, as you lose compile-time type checking.
When a value of a specific type is assigned to an interface{} variable, it gets automatically
boxed into the interface{} type This process is called type assertions.*/

package main

import "fmt"

func display[T any](x T) {
	fmt.Println(x)
}

func main() {
	display[string]("Focus on study")
	display[int](45)
	display[[]int]([]int{1, 2, 3, 4, 5})

	//omit type
	fmt.Println("Calling Generics in Omit type")

	display("Code is like humor. When you have to explain it, it’s bad.")
	display(23)
	display([]int{6, 7, 8, 9, 10})
}
