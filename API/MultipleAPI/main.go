package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/second", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Second")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}