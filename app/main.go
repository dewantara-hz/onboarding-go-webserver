package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Do something here
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":89", nil))
}
