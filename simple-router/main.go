package main

import (
	"fmt"
	"net/http"

	"github.com/kenlyau/goweb/simple-router/router"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foobar")
}
func main() {
	mux := router.NewRouter()
	mux.Add("/", helloHandler)
	mux.Add("/foo", fooHandler)
	http.ListenAndServe(":6060", mux)
}
