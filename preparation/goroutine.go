package main

import (
	"fmt"
	"sync"
	"time"
)

func things(name string) {
	fmt.Printf("Hello! %s \n", name)
	time.Sleep(time.Millisecond * 200)

}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {

		things("Sajid")
		wg.Done()
	}()
	go func() {
		things("Majid")
		wg.Done()
	}()

	wg.Wait()

}
