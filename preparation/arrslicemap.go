package main

import "fmt"

func main() {
	// array by using array literal
	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println("this is array by using array literal", arr)
	// simple array
	var arr1 [5]int

	arr1[0] = 101
	arr1[1] = 102
	arr1[2] = 103
	arr1[3] = 104
	arr1[4] = 105

	fmt.Println("This is a simple array", arr1)

	slice := []int{2, 5, 4, 6, 8, 9}

	fmt.Println("this is a simple slice :", slice)

	// by using make

	slice1 := make([]int, 5)

	fmt.Println("Slice using built in make: ", slice1)

	// simple map

	/*var m map[int]string

	  m[1] = "mohsin"
	  m[2] = "hamid"

	  fmt.Println("this is a simple map : ", m)*/

	// map literal

	map1 := map[int]string{
		1: "Sayyed",
		2: "Tauqeer",
	}

	fmt.Println("this is a map using map literal:", map1)

	// using built in make

	map2 := make(map[int]string)

	map2[1] = "Malik"
	map2[2] = "salman"

	fmt.Println("This is a map using make :", map2)

}
