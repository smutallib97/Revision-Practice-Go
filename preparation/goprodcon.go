package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	const workersCount = 3
	var buffer = make(chan int, workersCount)

	data := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("Stating Producer")

	go Producer(buffer, data)

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		fmt.Printf("Starting consumer #%d\n", i)
		go consumer(i, buffer, &wg)

	}
}
