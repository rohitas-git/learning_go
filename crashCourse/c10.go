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
// Defined Type

type meter float64
type km float64
type millimeter float64
type centimeter float64

func mToKm(dist meter) km{
	return km(dist/1000)
}

func (m meter) toKm() km{
	return km(m/1000)
}

func (m meter) toMilli() millimeter{
	return millimeter(m*1000)
}

func main() {
	var distance meter = 2999
	pl(mToKm(distance))

	pl(meter(10) > meter(9.999)) 
	
	pl(distance.toKm())
	pl(distance.toMilli())
}
