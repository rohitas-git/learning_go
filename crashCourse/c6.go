package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
// Command Line

func main() {
	// Print all the args using os package
	fmt.Println(os.Args)

	// Collect the args
	args := os.Args[1:]
	var iArgs = []int{}

	// Collect integer args from string args
	for _, v := range args {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}

	// Print the max int from iargs
	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}
	pl(max)

}
