package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(buff chan int, data []int) {
	defer close(buff)
	for _, i := range data {
		fmt.Printf("Producer sending %d to channel\n", i)
		buff <- i
	}
	fmt.Println("Producer exiting")
}

func consumer(idx int, buff chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		i, ok := <-buff
		if ok {
			fmt.Printf("Consumer #%d: received %d\n", idx, i)
			time.Sleep(1 * time.Second)
		} else {
			fmt.Printf("Consumer #%d: no more values to process, exiting\n", idx)
			return
		}
	}
}
func main() {
	var wg sync.WaitGroup
	const workersCount = 3
	var buffer = make(chan int, workersCount)
	data := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("Starting producer")
	go producer(buffer, data)

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		fmt.Printf("Starting consumer #%d\n", i)
		go consumer(i, buffer, &wg)
	}
	fmt.Println("Main waiting")
	wg.Wait()
	fmt.Println("Main exiting")
}
