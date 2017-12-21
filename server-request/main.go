package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./login.tpl")
		log.Println(t.Execute(w, nil))
	} else {
		var s string
		for k, v := range r.Form {
			s = fmt.Sprintf("%s\n%s:%s", s, k, v)
		}
		fmt.Fprintf(w, s)
	}
}
func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":6060", nil)
}
