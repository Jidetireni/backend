package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func greet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hello World! %s\n", p.ByName("name"))
}

func main() {

	mux := httprouter.New()
	mux.GET("/hello/:name", greet)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
