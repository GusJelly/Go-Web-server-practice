package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Home)

	log.Println("Started on port: ", portnum)
	fmt.Println("To close connection CTRL+C :-)")

	// Spinning up the server
	err := http.ListenAndServe(portnum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
