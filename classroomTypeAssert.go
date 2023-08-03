package main

import "fmt"

type Classroom interface {
	Study()
}

type Student struct {
	Name string
}

func (s Student) Study() {
	fmt.Println(s.Name, " is studying")
}
func main() {
	var classroom Classroom

	classroom = Student{"Sahil"}

	student1 := classroom.(Student)

	student1.Study()

}
