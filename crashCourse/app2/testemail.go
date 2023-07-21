package app2

import (
	"fmt"
	"regexp"
)

// For raw strings in Go: `` (backtick)
func IsEmail(s string) (string,error){
	r, _ := regexp.Compile(`[\w.%+-]{1,20}@[\w.-]{2,20}.[A-Za-z]{2,3}`) // regex for generic email

	if r.MatchString(s){
		return "Valid email", nil
	} else {
		return "",fmt.Errorf("Not a valid email")
	}
}