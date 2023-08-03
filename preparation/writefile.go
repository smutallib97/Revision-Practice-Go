package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "simple.txt"

	val := "old\nfalcon\nsky\ncup\forest\n"
	data := []byte(val)

	err := os.WriteFile(fileName, data, 0644)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
}
