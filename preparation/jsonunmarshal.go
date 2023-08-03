package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	EmpId  int
	Name   string
	City   string
	Salary int
}

func main() {
	jsoninput := `{
		"EmpId": 101,
		"Name": "Salim Khan",
		"City": "Malkapur",
		"Salary": 110000
	}`

	var employee Employee

	err := json.Unmarshal([]byte(jsoninput), &employee)

	if err != nil {
		fmt.Println("Json decode error..!")
	}
	fmt.Println(employee)

}
