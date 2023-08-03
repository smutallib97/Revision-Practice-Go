package main

import (
	"encoding/json"
	"fmt"
)

type friends struct {
	Name         string
	MobileNumber int
	City         string
}

func main() {
	friendDetails := []friends{
		{Name: "Tauqeer", MobileNumber: 8180883610, City: "Mumbai"},
		{Name: "Mayur", MobileNumber: 9518745632, City: "Nagpur"},
		{Name: "Sameer", MobileNumber: 7385485675, City: "Amravati"},
	}

	json_friendsBook, err := json.Marshal(friendDetails)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json_friendsBook))
}
