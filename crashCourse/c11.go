package main

import (
	"fmt"
	"log"
	"reflect"
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
// Interface

type Animal interface {
	// DrinkWater()
	HappySound()
	AngrySound()
}

type Cat string

func (c Cat) Attack() {
	pl("Cat attacks its prey")
}
func (c Cat) Name() string {
	return string(c)
}
func (c Cat) AngrySound() {
	pl("Hissssssss")
}
func (c Cat) HappySound() {
	pl("Purrrrrrrr")
}

func main() {
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()

	var kitty2 Cat = kitty.(Cat) // ?
	kitty2.Attack()
	pl("Cat name:", kitty2.Name())


}
