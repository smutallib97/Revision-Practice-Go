package main

import "fmt"

type Addressbook struct {
	Firstname string
	Lastname  string
	Address   string
	MobileNum string
	Email     string
}

func (a Addressbook) display() {
	fmt.Println(a.Firstname)
	fmt.Println(a.Lastname)
	fmt.Println(a.Address)
	fmt.Println(a.MobileNum)
	fmt.Println(a.Email)

}

func main() {
	addressbook1 := Addressbook{
		Firstname: "Sayyed",
		Lastname:  "Muzammil",
		Address:   "Taj Nagar Malkapur",
		MobileNum: "7815482583",
		Email:     "smuzz@gmail.com",
	}

	addressbook2 := Addressbook{
		Firstname: "Tauqeer",
		Lastname:  "Ahmad",
		Address:   "Taj Nagar Malkapur",
		MobileNum: "9563587889",
		Email:     "tauqeer@gmail.com",
	}

	addressbook1.display()
	addressbook2.display()

}
