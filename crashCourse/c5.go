package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

// create alias
var pl = fmt.Println

// UNUSED allows unused variables to be included in Go programs
func unused(x ...interface{}) {}
func newline() { pl("") }
func checkForErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/* ----------------------------------- TOC ---------------------------------- */
// File IO
// - Creating File
// - Reading File
// - Writing File
// - Check File exist
// - Set the Setting and Permission of File  
// Nice Error Handling

func main() {
	/* --------------------------------- File IO -------------------------------- */
	// Create the file
	f, err := os.Create("data.txt")
	checkForErr(err)
	defer f.Close() // close the file when out of main

	// Write series of prime in the file
	iPrimeArr := []int{2, 3, 5, 7, 11, 13}
	var sPrimeArr []string
	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")	// Write strings to file
		checkForErr(err)
	}

	// Open the file
	f, err = os.Open("data.txt")
	checkForErr(err)
	defer f.Close() // close the file when out of main

	// Read the file
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("Prime:", scan1.Text())
	}
	// Error handling for reading
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}

	// Check if the file exist and 
	_, err = os.Stat("data.txt")

	// Another way of error handling; here catching a specific error
	if errors.Is(err, os.ErrNotExist) {
		pl("File doesn't exist")
	} else {
		// Set the file to be appended or created or write-only and decide its permissions (0644)
		f,err = os.OpenFile("data.txt",
			os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
		checkForErr(err) 
	}
	defer f.Close()

	// Write string to file
	if _, err := f.WriteString("17\n"); err !=nil {
		log.Fatal(err)
	}

}
