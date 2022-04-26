package main

import (
	"fmt"
	"net/http"
)

func main() {

	myHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	}

	myCatHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello I'm a cat")
	}

	http.HandleFunc("/", myHandler)
	http.HandleFunc("/cat", myCatHandler)

	http.ListenAndServe(":8000", nil)
}
