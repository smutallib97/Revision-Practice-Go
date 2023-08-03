package main

import "fmt"

type Animals interface {
	Sound() string
}

type Dog struct {
}

type Cat struct {
}

func (c Cat) Sound() string {
	return "Meow"
}
func (d Dog) Sound() string {
	return "Woof"
}

func main() {

	animals := []Animals{Dog{}, Cat{}}

	for _, animal := range animals {
		fmt.Println(animal.Sound())
	}

}
