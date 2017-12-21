package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/foo/bar")
}
func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/foo", fooHandler)
	http.ListenAndServe(":6060", nil)
}
