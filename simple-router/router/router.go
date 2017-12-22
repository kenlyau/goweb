package router

import (
	"fmt"
	"net/http"
)

type Handle func(http.ResponseWriter, *http.Request)

type Router struct {
	mux map[string]Handle
}

func NewRouter() *Router {
	return &Router{
		mux: make(map[string]Handle),
	}
}
func (router *Router) Add(path string, handle Handle) {
	router.mux[path] = handle
}
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(r.URL.Path)
	for path, handler := range router.mux {
		if r.URL.Path == path {
			handler(w, r)
			return
		}
	}
	http.NotFound(w, r)
	return
}
