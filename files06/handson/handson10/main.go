package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl01, tpl02, tpl03, tpl04, tpl05 *template.Template

func init() {
	tpl01 = template.Must(template.ParseFiles("templates/index.gohtml"))
	tpl02 = template.Must(template.ParseFiles("templates/apply.gohtml"))
	tpl03 = template.Must(template.ParseFiles("templates/applyProcess.gohtml"))
	tpl04 = template.Must(template.ParseFiles("templates/about.gohtml"))
	tpl05 = template.Must(template.ParseFiles("templates/contact.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/apply", apply)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl01.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request) {
	err := tpl04.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request) {
	err := tpl05.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		err := tpl03.ExecuteTemplate(w, "applyProcess.gohtml", nil)
		HandleError(w, err)
		return
	}
	err := tpl02.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
