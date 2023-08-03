package main

import "fmt"

func main() {

	for i, j := range "Sayyed" {
		fmt.Printf("the index number of %U is %d \n", j, i)
	}
}
