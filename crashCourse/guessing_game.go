package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// create alias
var pl = fmt.Println

// UNUSED allows unused variables to be included in Go programs
func unused(x ...interface{}) {}
func newline(){pl("")}

/* ----------------------------------- TOC ---------------------------------- */
//

func main(){
	seedSec := time.Now().Unix()
	rand.Seed(seedSec)
	randNum := rand.Intn(50) + 1

	for true {
		fmt.Println("Guess a number between 0 and 50")
		// pl("Random number:", randNum) 

		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}

		if iGuess > randNum {
			pl("Guess Lower")
		} else if iGuess < randNum{
			pl("Guess Higher")
		} else {
			pl("Guess is Equal")
			break
		}

	}
}