package main

import "fmt"

func main() {
	var x int
	x = 42
	ptr := &x // referencing

	value := *ptr //dereferencing

	fmt.Println(ptr)
	fmt.Println(value)

}
