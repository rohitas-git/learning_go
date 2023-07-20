package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

// create alias
var pl = fmt.Println

// UNUSED allows unused variables to be included in Go programs
func unused(x ...interface{}) {}

// Variables
// Data types
// Casting

func main() {

	pl("What is your name?")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}

	// var name type
	var vName string = "Rohi"
	var v1, v2 = 1.2, 1.4
	var v3 = "hello"
	v4 :=  2.0
	v4 = 5.0
	unused(v1,v2,v3,v4,vName)

	// data types
	// int, float64, bool, string, rune, 
	// Default value: 0, 0.0, false, ""
	pl(reflect.TypeOf(25))	
	pl(reflect.TypeOf(2.0))	
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("hello"))	

	// type casting
	cV1 := 2.5
	cV2 := int(cV1)
	pl(cV2)

	// string --> int
	cv3 := "1235"
	cv4, err := strconv.Atoi(cv3)
	pl(cv4, err, reflect.TypeOf(cv4))

	// int --> ASCII/ string
	cv5 := 1235
	cv6 := strconv.Itoa(cv5)
	pl(cv6, reflect.TypeOf(cv6))

	// string --> float64
	cv7 := "3.14"
	if cv8, err := strconv.ParseFloat(cv7, 64); err == nil {
		pl(cv8, reflect.TypeOf(cv8))
	}

	// float --> string
	cv9 := fmt.Sprintf("%f", 3.14)
	pl(cv9, reflect.TypeOf(cv9))
}

