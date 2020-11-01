package main

import (
	"fmt"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", hellohandler)
	http.ListenAndServe(":8888", nil)
}
