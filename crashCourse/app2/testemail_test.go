package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T){
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("hello is not an email")
	}	
	_, err = IsEmail("ola@mail.com")
	if err == nil {
		t.Error("ola@mail.com is an email")
	}
	_, err = IsEmail("ola@mail")
	if err != nil {
		t.Error("ola@mail is an email")
	}
}