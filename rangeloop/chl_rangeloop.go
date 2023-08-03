package main

import "fmt"

func main() {
	chl := make(chan string)

	go func() {
		chl <- "Sayyed"
		chl <- "Muzammil"
		chl <- "Mayur"

		close(chl)
	}()

	for i := range chl {
		fmt.Println(i)
	}
}
