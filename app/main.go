package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/dewantara-hz/onboarding/reqres"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Homepage
	fmt.Fprintf(w, "Hi there, welcome!")
}

func validationHandler(w http.ResponseWriter, r *http.Request) {
	url, _ := r.URL.Query()["url"]
	valType, _ := r.URL.Query()["type"]

	body := reqres.Access(url[0])
	var validity bool

	if valType[0] == "json" {
		validity = reqres.IsJson(body)
	} else if valType[0] == "xml" {
		validity = reqres.IsXML(body)
	} else if valType[0] == "html" {
		validity = reqres.IsHTML(body)
	}

	fmt.Fprint(w, validity)

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/validator/", validationHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
