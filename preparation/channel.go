package main

import "fmt"

func main() {
	mychnl := make(chan string)

	go func() {
		mychnl <- "Accept"
		mychnl <- "Your"
		mychnl <- "Fear"
		close(mychnl)
	}()

	for res := range mychnl {
		fmt.Println(res)
	}
}
