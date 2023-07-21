package main

import (
	"fmt"
	"log"
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
// Packages (Library) --> ./app

// ./ $ go mod init <currentRefName>
// 
// └── app
//     ├── go.mod
//     ├── main.go
//     └── mypackage
//         └── mypackage.go
// 
// In mypackage.go
// 		package <toImportPackageName>
// 		var ToExposeVar int
// 		func ToExposeFunc(params) {BODY}
// 
// 
// In main.go
// 		import <toImportPackageName> "<currentRefName/mypackage>"

