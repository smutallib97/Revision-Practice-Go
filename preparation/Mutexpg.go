package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type User struct {
	name           string
	depositAmount  float32
	withdrawAmount float32
}

var BALANCE float32 = 12000

func (u *User) deposit(wg *sync.WaitGroup, mx *sync.Mutex) {
	fmt.Printf("%s is depositing  rupees %f \n", u.name, u.depositAmount)
	mx.Lock()
	BALANCE += u.depositAmount
	defer mx.Unlock()
	wg.Done()
}

func (u *User) withdraw(wg *sync.WaitGroup, mx *sync.Mutex) {
	fmt.Printf("%s is withdrawing  rupees %f \n", u.name, u.withdrawAmount)
	mx.Lock()
	BALANCE -= u.withdrawAmount
	defer mx.Unlock()
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex
	users := []User{
		{name: "Ammar Khan", withdrawAmount: 1300, depositAmount: 1000},
		{name: "Ramesh Patel", withdrawAmount: 1400, depositAmount: 123},
		{name: "Sahil Satpute", withdrawAmount: 900, depositAmount: 25},
		{name: "Sunil Chavan", withdrawAmount: 200, depositAmount: 5432},
		{name: "Mayur Polojwar", withdrawAmount: 5600, depositAmount: 2344},
	}
	rand.Seed(time.Now().UnixNano())
	for i := range users {
		wg.Add(2)
		i = rand.Intn(len(users))
		go users[i].deposit(&wg, &mu)
		go users[i].withdraw(&wg, &mu)
		time.Sleep(time.Second)
	}

	wg.Wait()

	fmt.Printf("New account BALANCE is %f \n", BALANCE)

}
