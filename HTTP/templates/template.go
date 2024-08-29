package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "Hello world")

	// 	tmpl := `<!DOCTYPE html>
	// <html>
	//  <head>
	//  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	//  <title>Go Web Programming</title>
	//  </head>
	//  <body>
	//  {{ . }}
	//  </body>
	// </html>`

	// 	t := template.New("tmpl.html")
	//  t, _ = t.Parse(tmpl)

	// t, err := template.ParseFiles("tmpl.html")
	// if err != nil {
	//     // Handle the error
	// }
	// t = template.Must(template.ParseFiles("tmpl.html"))

	// t, _ := template.ParseFiles("t1.html", "t2.html")
	// t.ExecuteTemplate(w, "t2.html", "Hello World!")

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
