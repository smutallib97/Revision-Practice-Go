package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("quote.txt")

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exits")
	} else {
		fmt.Println("File exits")
	}

}
