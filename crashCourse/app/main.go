package main

import (
	stuff "example/project/mypackage"
	"fmt"
	"log"
)

func main() {
	fmt.Println("My name is", stuff.Name)
	intArr := []int{2,3,5,7,11}
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Println(strArr)

	// How encapsulation looks like
	date := stuff.Date{}
	err := date.SetDay(15)	// Using Setter func
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(8)	// Using Setter func
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetYear(1947)	// Using Setter func
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Day(),date.Month(),date.Year())
}