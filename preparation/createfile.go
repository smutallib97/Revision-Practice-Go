package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("simple.txt")

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file created")
}
