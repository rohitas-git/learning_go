package main

import (
	"fmt"
	"log"
	"regexp"
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
// Regular Expression

func main() {
	reStr := "The ape was at the apex"
	match, _ := regexp.MatchString("(ape[^ ]?)", reStr) // ape that is not followed by space
	pl(match)

	reStr2 := "Cat rat mat fat pat"
	r,_ := regexp.Compile("([crmfp]at)")
	pl(r.MatchString(reStr2))
	pl(r.FindString(reStr2))
	pl(r.FindStringIndex(reStr2))
	pl(r.FindAllString(reStr2, -1))
	pl(r.FindAllString(reStr2, 2))
	pl(r.FindAllStringSubmatchIndex(reStr2, -1))
	pl(r.ReplaceAllString(reStr2, "Dog"))

}

