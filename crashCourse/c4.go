package main

import (
	"fmt"
)

// create alias
var pl = fmt.Println

// UNUSED allows unused variables to be included in Go programs
func unused(x ...interface{}) {}
func newline(){pl("")}

/* ----------------------------------- TOC ---------------------------------- */
// Functions
// Pointers


func getSq(x int, y int) (int, int){
	return x*x, y*y
}

func getQuotient(x float64, y float64) (ans float64, err error){
	if y == 0 {
		return 0, fmt.Errorf("You can't divide by zero")
	} else {
		return x/y , nil
	}
}

func getSum(nums ...int) int{
	sum :=0
	for _, num := range nums{
		sum +=num
	}
	return sum
}

func getArraySum(arr []int) int{
	sum :=0
	for _, num := range arr{
		sum +=num
	}
	return sum
}

// pass by value
func tryChangeVal(f3 int) int{
	f3 += 1
	return f3
}

// pass by reference
func tryChangeVal2(f3Ptr *int) {
	*f3Ptr = 121
}

func getAverage(nums ...float64) float64{
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, val := range nums{
		sum += val
	}
	return (sum/numSize)
}

func main(){

	/* -------------------------------- Functions ------------------------------- */
	// func funcName (params) returnType {BODY}

	pl(getSq(12,23))

	pl(getQuotient(5,0))
	pl(getQuotient(13,11))

	pl(getSum(1,2,3,4,5,6,7,8,9))


	vArr := []int{10,11,12,13,14,15,16}
	pl(getArraySum(vArr))

	f3:= 5
	pl("f3 before:", f3)	// 5
	tryChangeVal(f3)
	pl("f3 after:", f3) 	// 5

	f3 = 10
	pl("f3 before:", f3)	// 10
	tryChangeVal2(&f3)
	pl("f3 after:", f3) 	// 121

	/* --------------------------------- Pointer -------------------------------- */
	f4 := 33
	var f4Ptr *int = &f4
	pl("Address of f4:", f4Ptr)
	pl("Value of f4:", *f4Ptr)

	iSlice := []float64{112,13,51}
	fmt.Printf("Average : %.3f\n", getAverage(iSlice...))
}