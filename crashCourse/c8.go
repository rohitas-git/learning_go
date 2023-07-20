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
// Maps

func main() {
	/* --------------------------------- Mapping -------------------------------- */
	// Collection of key-value pairs
	// keys can be any datatype that can be compared using == 

	// var myMap  map [keyType]valueType

	var heroes map[string]string
	heroes = make(map[string]string)
	villains := make(map[string]string)

	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["Spiderman"] = "Peter Parker"

	villains["Lex Luther"] = "Bald Brooder"
	villains["Thanos"] = "Purple Potato"

	superPets := map[int]string{ 1: "Krypto", 2: "Bat Hound"}
	unused(superPets)

	fmt.Printf("Batman is %v\n", heroes["Batman"])
	pl("Chip:", superPets[3])
	_, ok := superPets[3]
	pl("Is there a 3rd pet: ", ok)

	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k,v)
	}

	// delete entry for specific key from the map
	delete(heroes, "Spiderman")
}
