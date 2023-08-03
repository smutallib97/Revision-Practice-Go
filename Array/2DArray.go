package main

import "fmt"

func main() {

	var twoarr [3][2]int

	twoarr[0][0] = 10
	twoarr[1][0] = 20
	twoarr[2][0] = 30

	for i := 0; i < len(twoarr); i++ {
		twoarr[i][1] = twoarr[i][0] * 2
	}
	fmt.Println(twoarr)
}
