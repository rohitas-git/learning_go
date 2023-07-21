package main

import (
	"fmt"
	"log"
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
// Concurrency / Go Routines


/* ------------------------- Random Order Execution ------------------------- */
func printTo15() {
	for i := 1; i < 15; i++ {
		pl("Fun 1: ", i)
	}
}
func printTo10() {
	for i := 1; i < 10; i++ {
		pl("Fun 2: ", i)
	}
}

// Cannot now the order of execution of printTo15 and printTo10
// Non-determinstic
func main() {
	go printTo10()
	go printTo15()
	time.Sleep(2 * time.Second) // so that main doesn't immediately exits
}
