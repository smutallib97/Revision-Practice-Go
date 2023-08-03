package main

import (
	"fmt"
)

func main() {

	rmap := map[int]string{
		11: "Tauqeer",
		12: "Muzammil",
		13: "Junaid",
	}

	for key, value := range rmap {
		fmt.Println(key, value)
	}
}
