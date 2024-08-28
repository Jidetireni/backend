package main

import (
	"fmt"
	"net/http"
)

// func setCookie(w http.ResponseWriter, r *http.Request) {
// 	c1 := http.Cookie{
// 		Name:     "first_cookie",
// 		Value:    "Go web Programming",
// 		HttpOnly: true,
// 	}
// 	c2 := http.Cookie{
// 		Name:     "second_cookie",
// 		Value:    "banm Publication co",
// 		HttpOnly: true,
// 	}
// 	w.Header().Set("Set-Cookie", c1.String())
// 	w.Header().Set("set-cookie", c2.String())
// }

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "banm Publication co",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)

}

func getcookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "cannot get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getcookie)
	server.ListenAndServe()
}
