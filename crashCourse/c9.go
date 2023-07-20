package main

import (
	"fmt"
	"log"
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
// Generics

func main() {

}
