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
// Generics and Constraints
// Structs
// Composition

type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T{
	return x + y
}

type customer struct {
	name, address string
	bal float64
}

type contact struct {
	phone int
	customer
}

func printCustomerInfo(this customer){
	pl("Name: ", this.name)
	pl("Address: ", this.address)
	pl("Balance: ", this.bal)
}

func newCustomerAddress(c *customer, d string) {
	c.address = d
}

func (c contact) address(){
	pl(c.customer.address)
}
func (c contact) phoneNumber(){
	pl(c.phone)
}

func main() {
	/* -------------------------------- Generics -------------------------------- */
	pl(getSumGen(10,5))
	pl(getSumGen(4.5,5.9))

	/* --------------------------------- Structs -------------------------------- */

	var tS customer
	tS.name = "Sherlock Holmes"
	tS.address = "London 221"
	tS.bal = 10000

	newCustomerAddress(&tS, "221B Baker Street")
	printCustomerInfo(tS)

	another := customer{"Bruce Wayne (Batman)", "Wayne Manor, Gotham City, NY, USA", 0.0}
	printCustomerInfo(another)

	// ! Go doesn't support Inhertiance
	// ! But supports composition
	/* ------------------------------- Composition ------------------------------ */

	// customer struct is embedded into contact struct 
	aContact := contact{99247897, another}
	aContact.address()
	aContact.phoneNumber()

}
