package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, res *http.Request) {
	io.WriteString(w, "This is my index page!!")
}

func dogHandler(w http.ResponseWriter, res *http.Request) {
	io.WriteString(w, "DOG is DOG")
}

func meHandler(w http.ResponseWriter, res *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(w, "something.gohtml", "Noboru Nakahara")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(indexHandler))
	http.Handle("/dog/", http.HandlerFunc(dogHandler))
	http.Handle("/me/", http.HandlerFunc(meHandler))

	http.ListenAndServe(":8080", nil)
}
