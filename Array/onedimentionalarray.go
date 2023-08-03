//array is a homogenous data structure with fixed length.

package main

import "fmt"

func main() {
	var arr_1D [5]int

	//var i, j int

	for i := 0; i < 5; i++ {
		arr_1D[i] = i
	}
	for j := 0; j < 5; j++ {
		fmt.Printf("Element (%d) = %d \n", j, arr_1D[j])
	}
}
