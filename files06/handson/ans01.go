package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/sample.png", dogPic)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset = utf-8")
	io.WriteString(w, `
	foo ran
	`)
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("dog.gohtml"))
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "sample.png")
}
