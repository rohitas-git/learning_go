package main

import (
	"fmt"
	stuff "example/project/mypackage"
)

func main() {
	fmt.Println("My name is", stuff.Name)
	intArr := []int{2,3,5,7,11}
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Println(strArr)
}