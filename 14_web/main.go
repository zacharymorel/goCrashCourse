package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1>Hello World!</H1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1>ABOUT!</H1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Runing")
	http.ListenAndServe(":3000", nil)
}
