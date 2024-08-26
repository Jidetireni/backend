package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, "(1)", r.Form)
	fmt.Fprintln(w, "(2)", r.PostForm)
	fmt.Fprintln(w, "(3)", r.FormValue("hello"))
	fmt.Fprintln(w, "(4)", r.PostFormValue("hello"))
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, "(5)", r.MultipartForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
