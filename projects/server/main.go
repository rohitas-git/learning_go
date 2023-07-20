package main

import(
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	
}