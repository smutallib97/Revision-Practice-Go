/*Write a function called sumSlice that takes an interface{} as an argument.
The interface{} can be any slice containing numeric values (int, float64, etc.).

it means

slice := []int{1,2,3,4,5}
slcie1 :=[]string {"stringone", "stringtwo"}
slice2 := []float64{10.2,45.6,65.6,77.22,89.0}


1+2+3+4+5+6+7+8+9

1.0+2.3+4.6+7.8+9.9

"one"+"two"+"three"+"four"

The function should use a ===type assertion === to extract the values from the slice, sum them up,
and return the total sum as a float64.*/

package main

import "fmt"

func sumSlice(inputData interface{}) float64 {

	slice, ok := inputData.([]interface{})

	if !ok {
		// if its not a slice, return 0
		return 0
	}

	var total float64

	for _, value := range slice {
		switch v := value.(type) {
		case int:
			total += float64(v)
		case float64:
			total += v
		case float32:
			total += float64(v)
		default:
			total = 0

		}
	}

	if total == 0 {
		fmt.Println("Non supported slice")
	}
	return total

}

func main() {
	nums := []interface{}{1, 2, 3, 4, 5}
	sum := sumSlice(nums)
	fmt.Println(sum) // 15

	nums2 := []interface{}{"string1", "string2", "string3"}
	sum2 := sumSlice(nums2)
	fmt.Println(sum2) // 15

	num3 := [3]int{1, 3, 4}
	sum3 := sumSlice(num3)
	fmt.Println(sum3)
}
