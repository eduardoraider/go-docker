package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(rw, "Hello World")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
