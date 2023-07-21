package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// create alias
var pl = fmt.Println

// UNUSED allows unused variables to be included in Go programs
func unused(x ...interface{}) {}
func newline()                { pl("") }
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/* ----------------------------------- TOC ---------------------------------- */
// Lock and Mutex

type Account struct {
	balance int
	lock    sync.Mutex // allow only one customer to access this account at one time
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) WithdrawBalance(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.balance < v {
		pl("Not enough balance")
	} else {
		fmt.Printf("Withdraw %d from balance %d\n", v, a.balance)
	}
	a.balance -= v
}

func main() {
	var acc Account
	acc.balance = 100
	pl("balance of account:", acc.GetBalance())

	for i := 0; i < 12; i++ {
		go acc.WithdrawBalance(10)
	}

	time.Sleep(2 * time.Second)
}
