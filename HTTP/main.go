package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Post struct {
	User    string
	Threads []string
}

func requestHeaders(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rh := r.Header
	fmt.Fprintln(w, rh)
}

func requestBody(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lenght := r.ContentLength
	body := make([]byte, lenght)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func writeResponse(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

func writeStatusCode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func responseHeader(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("location", "http://google.com")
	w.WriteHeader(302)
}

func jsonResponse(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	post := Post{
		User:    "jide-jegede",
		Threads: []string{"first", "second", "third"},
	}
	jsonData, _ := json.Marshal(post)
	w.Write(jsonData)
}

func main() {
	router := httprouter.New()
	router.GET("/request_headers", requestHeaders)
	router.GET("/request_body", requestBody)
	router.GET("/write_reponse", writeResponse)
	router.GET("/write_statuscode", writeStatusCode)
	router.GET("/redirct", responseHeader)
	router.GET("/json", jsonResponse)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	server.ListenAndServe()
}
