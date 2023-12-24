package main

import (
	"fmt"
	"net/http"
)

const portnum string = ":8080"

// Handler function
func Home(w http.ResponseWriter, r *http.Request) {
	// Message here to the response writer
	fmt.Fprintf(w, "Welcome to my home!")
}
