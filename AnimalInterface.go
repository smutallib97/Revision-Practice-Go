//Interface

// It is set of one or more methods signature 
//It is like a blue print of any data type can follow  if it wants to be treated as that interface type

package main

import "fmt"

type Animal interface{
	makeSound() string
	move() string
}

type Dog struct{

}
func (d Dog ) makeSound() string{
	return "Woof!"
}

func (d Dog) move() string{
	return "Dog can walk"
}
func main(){

	var animal Animal

	animal = Dog{}
	fmt.Println(animal.makeSound())
	fmt.Println(animal.move())

}