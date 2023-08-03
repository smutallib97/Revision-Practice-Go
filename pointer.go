package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0 // dereferencing
}
func main() {
	x := 5
	zero(&x) // referencing

	fmt.Println(x)

}
