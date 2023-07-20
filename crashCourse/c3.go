package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// create alias
var pl = fmt.Println

// UNUSED allows unused variables to be included in Go programs
func unused(x ...interface{}) {}
func newline(){pl("")}

/* ----------------------------------- TOC ---------------------------------- */
// Random Number, Seeding with Time
// Maths functions
// Formatted String
// For Loops 
// Arrays
// Slices


func main(){
	/* ------------------------------- Random Seed ------------------------------ */
	seedSec := time.Now().Unix()
	// pl(seedSec)
	rand.Seed(seedSec)
	randNum := rand.Intn(50) +1
	pl("Random: ", randNum)

	/* ---------------------------------- Math ---------------------------------- */
	// math.<_>(...): 
	// - Abs, Pow, Sqrt, Cbrt, Ceil, Floor, Round, Log2, Log10, Log, Max, Min, Exp
	// - Sin, Cos, ...

	/* ----------------------------- Formatted Print ---------------------------- */
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied data type

	fmt.Printf("%s, %d, %c, %f, %t, %o, %x\n", "Stuff", 12, 'A', 3.123, true, 1, 1)
	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.141592)
	fmt.Printf("%9.2f\n", 3.141592)
	fmt.Printf("%9.f\n", 3.141592)

	sv1 := fmt.Sprintf("%9.f ", 3.141692)
	pl(sv1, reflect.TypeOf(sv1))

	/* ---------------------------------- LOOP ---------------------------------- */
	// for init; condn; postStatement {Body}
	for x :=1; x<5; x++ {
		pl(x)
	}
	newline()

	// while loop
	fX := 5
	for fX > 0 {
		pl(fX)
		fX--
	}

	/* -------------------------------- ARRAY -------------------------------- */
	aNums := []int{11,22,33}
	for i, num := range aNums {
		fmt.Printf("%d: %d\n", i, num)
	}

	var arr1 [5]int
	arr1[0] = 1
	arr2 := [5]int{1,2,99}
	pl(arr2, len(arr2))

	arr3 := [3][3]int {
		{1,2,3},
		{2,3,1},
		{3,1,2},
	}
	for x:=0; x < 3; x++{
		pl(arr3[x])
	}

	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr{
		fmt.Printf("Rune array: %d\n", v)
	}

	// byte array --> string
	byteArr := []byte{'r','o','h','i'}
	bStr := string(byteArr[:])
	pl(bStr)

	/* ---------------------------------- Slice --------------------------------- */
	// like array but can grow
	
	// var name []datatype
	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Gods"
	pl(len(sl1))
	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])	
	}

	sArr := [5]int{1,2,3,4,5} 
	pl(sArr[:])
	pl(sArr[:3])
	pl(sArr[3:])

	// Slice and Array affect each other
	sl2 := sArr[:3]
	sArr[0] = 10
	pl(sArr, sl2)
	sl2[0] = 1
	pl(sArr, sl2)
	
	sl3 := append(sl2,12)
	pl("sl3:", sl3)
	pl("sArr:", sArr)

	sl4 := make([]string, 5)
	pl("sl4:", sl4," type:", reflect.TypeOf(sl4))
	pl("sl4[0]:", sl4[0], " type:", reflect.TypeOf(sl4[0]))
	rsl := []rune(sl4[0])
	pl(rsl)
}