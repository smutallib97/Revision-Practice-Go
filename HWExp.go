/*Create a struct called Person with the following fields: Name (string) and Age (int).
Now, write a function called printPersonDetails that takes an interface{} as an argument.
The interface{} can be either a *Person or a string. Use a type assertion to convert the interface{}
to the appropriate type and then print the person's name and age if it's a *Person,
 or print the string directly if it's a string. */

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func printPersonDetails(inputData interface{}) {

	/*person, ok := inputData.(*Person)
	str, ok := inputData.(string)

	if ok {
		fmt.Printf("Name: %s", person.Name)
		fmt.Printf("Age:%d", person.Age)
	} else if ok {
		fmt.Println(str)
	} else {
		fmt.Println("Invalid Input Data")
	}*/

	switch value := inputData.(type) {
	case *Person:
		fmt.Printf("Name : %s, Age: %d \n\n", value.Name, value.Age)

	case string:
		fmt.Println(value)

	default:
		fmt.Println("Invalid Data Input")
	}

}

func main() {
	//for person details
	person1 := &Person{Name: "Tauqeer", Age: 26}
	printPersonDetails(person1)
	// for string
	printPersonDetails("Hello Bridgelabians")
	// if there any other type of data
	printPersonDetails(45)

}
