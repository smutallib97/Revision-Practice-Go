package main

import (
	"fmt"
	"time"
)

func EmpName() {
	empNames := []string{"Sayyed", "Tauqeer", "Ramesh", "Mayur"}

	for t1 := 0; t1 <= 3; t1++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%s \n", empNames[t1])
	}
}

func EmpId() {
	empIds := []int{101, 102, 103, 104}

	for t2 := 0; t2 <= 3; t2++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d \n", empIds[t2])
	}
}

func main() {
	fmt.Println("Goroutine Started....!!!")

	go EmpName()

	go EmpId()

	time.Sleep(3500 * time.Millisecond)
	fmt.Println("\n!...Main Go-routine End...!")
}
