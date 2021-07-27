package main

import (
	"fmt"
	"net/http"
)

//Handler Function
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Web World</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About Us</h1>")
}

func main() {

	//Our Routes
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}
