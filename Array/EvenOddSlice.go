package main

import "fmt"

func isEven(arr []int) {
	var evenSlice []int

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			evenSlice = append(evenSlice, arr[i])
		}
	}
	fmt.Println("Even Elements Slice is :", evenSlice)
}

func isOdd(arr []int) {
	var oddSlice []int

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			oddSlice = append(oddSlice, arr[i])
		}
	}
	fmt.Println("Odd Slice Elements are: ", oddSlice)
}

func main() {

	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println("20 Interger numbers array is : ", arr1)

	slice1 := arr1[5:15]
	fmt.Println("10 elements from the original array: ", slice1)

	isEven(arr1) // calling the even function for priting the evenSlice
	isOdd(arr1)  //calling the odd function for priting the oddSlice
}
