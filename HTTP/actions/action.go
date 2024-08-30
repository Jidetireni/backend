package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func process2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl2.html")
	daysOfWeek := []string{"mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, daysOfWeek)
}

func process3(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl3.html")
	t.Execute(w, "hello")
}

func process4(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.Execute(w, "Hello World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/process2", process2)
	http.HandleFunc("/process3", process3)
	http.HandleFunc("/process4", process4)
	server.ListenAndServe()
}
