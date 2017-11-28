package main

import (
	"fmt"
	"log"
	"net/http"
)

func callHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello old world!")
}

func main() {
	http.HandleFunc("/", callHelloWorld)
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
