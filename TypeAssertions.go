package main

import "fmt"

type Bird interface {
	fly()
}
type Parrot struct {
	Name string
}
type Sparrow struct {
	Name string
}

type Eagle struct {
	Name string
}

// implement the method of the interface then only it will become bird
func (p Parrot) fly() {
	fmt.Println(p.Name, "bird is flying")
}

func (s Sparrow) fly() {
	fmt.Println(s.Name, "Bird is flying")

}

func (e Eagle) fly() {
	fmt.Println(e.Name, "Bird is flying")

}
func main() {
	var bird Bird

	bird = Parrot{"Shikha"}

	// We will write 300 lines of code

	bird = Eagle{"Toofan"}

	//We will write 300 lines of code
	// now i am confued which bird is there in the bird object,
	//because i have written 300 and 400 lines of code

	// now i use type assertion to check bird type

	eagle := bird.(Eagle)

	eagle.fly()

	parrot := bird.(Parrot)
	parrot.fly()
}
