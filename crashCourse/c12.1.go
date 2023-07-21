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
// Channels


/* ------------------------- Fixed Order Execution ------------------------- */
// Func calls that fill up the argument channel; channel <- ValFilledType
// In main:
// make a channel; 	channel's type: chan <Filledtype>
// fill that channel using above written func;	fillFunc(channelName)
// print out each item filled in channel;  ValFilledType from (<- channelName)

func nums1(channel chan int){
	channel <- 1
	channel <- 2
	channel <- 3
}

func nums2(channel chan int){
	channel <- 4
	channel <- 5
	channel <- 6
}

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	go nums1(channel1)
	go nums2(channel2)

	pl(<- channel1)
	pl(<- channel1)
	pl(<- channel1)
	pl(<- channel2)
	pl(<- channel2)
	pl(<- channel2)

}
