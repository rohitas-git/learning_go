package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

// create alias
var pl = fmt.Println

// UNUSED allows unused variables to be included in Go programs
func unused(x ...interface{}) {}
func newline(){pl("")}

/* ----------------------------------- TOC ---------------------------------- */
// Operators
// Strings
// Runes
// Time

func main(){
	// Conditional Operators: > < >= <= == !=
	// Logical Operators: && || !
	// Easy stuff; left it

	// String: array of bytes - []byte

	sv1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sv2 := replacer.Replace(sv1)
	pl(sv1, "-->", sv2)
	pl("Length:", len(sv2))
	pl("Contains 'Another' :", strings.Contains(sv2, "Another"))
	pl("'o' index:", strings.Index(sv2, "o"))
	pl("Replace:", strings.Replace(sv2,"o","0", -1))
	sv3 := "\nSome Words\n"
	// pl(sv3)
	sv4 := strings.TrimSpace(sv3)
	pl("Trimed from \\nSome Words\\n: ", sv4)
	pl("Split:", strings.Split("a, b, c, d", ", "))
	pl("Lower:", strings.ToLower(sv2))
	pl("Upper:", strings.ToUpper(sv2))
	pl("Prefix:", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix:", strings.HasSuffix("tacocat", "cat"))
	newline()

	// Character (aka Runes)
	rStr := "abcdefg"
	pl("Rune count:", utf8.RuneCountInString(rStr))

	for i, runeVal := range rStr {
		// pl(i, ":", runeVal)
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}

	// Time
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())
	pl(now.Date())
}