package main

import (
	"fmt"
	"net/http"
)

type Mux struct{}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		helloHandler(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
func main() {
	mux := &Mux{}
	http.ListenAndServe(":6060", mux)
}
