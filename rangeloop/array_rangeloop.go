package main

import (
	"fmt"
)

func main() {
	//declaration of name array
	names := []string{"Sahil", "Tauqeer", "Sayyed", "Muzammil"}

	//i stores index no. of individuals
	//j stors individual value of string

	for i, j := range names {
		fmt.Println(i, j)
	}
}
