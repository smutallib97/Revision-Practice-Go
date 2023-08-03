package main

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	Name                       string
	Period                     string
	Matches_Played, RunsScored int
	AvgScore                   float64
}

func main() {
	fmt.Println("Print the variable of type Player: ")

	player1 := Player{"Sachin Tendulkar", "1989-2012", 463, 18426, 44.83}
	player2 := Player{"Saurav Ganguly", "1992-2007", 308, 11221, 40.95}
	player3 := Player{"Rahul Dravid", "1996-2011", 340, 10768, 39.15}

	players := []Player{
		player1,
		player2,
		player3,
	}
	for _, plyr := range players {
		fmt.Println(plyr)
	}
	fmt.Println(player1)

	fmt.Println("")
	fmt.Println("Converting variable to json")
	bytes, _ := json.Marshal(player1)
	fmt.Printf("%T %s", bytes, string(bytes))
}
