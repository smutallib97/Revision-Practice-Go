package main

import "fmt"

type Friendbook struct {
	Name         string
	Address      string
	City         string
	MobileNumber string
}

func (f Friendbook) displayFriendbook() {
	fmt.Println("Name: ", f.Name)
	fmt.Println("Address:", f.Address)
	fmt.Println("City:", f.City)
	fmt.Println("MobileNumber:", f.MobileNumber)
}
func main() {
	Friend1 := Friendbook{
		Name:         "Tauqeer Ahamed",
		Address:      "Santosh Nagar",
		City:         "Mumbai",
		MobileNumber: "8888123610",
	}

	Friend1.displayFriendbook()

}
