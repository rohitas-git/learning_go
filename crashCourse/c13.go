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
// Closures
// Pass a function to a function
// Recursion

func useFunc(f func(int, int) int, x int, y int) {
	pl("Answer:", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	/* --------------------------------- Closure -------------------------------- */
	intSum := func(x, y int) int { return x + y }
	pl(intSum(11, 3))

	samp1 := 1
	changeVar := func() { samp1++ }

	changeVar() // Closure can change values outside of function
	pl(samp1)   // Now samp1 =  2

	/* --------------------------- Passing fn to a fn --------------------------- */
	useFunc(sumVals, 5, 8)

	/* -------------------------------- Recursion ------------------------------- */
	pl(factorial(4))
}
