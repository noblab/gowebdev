package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", block)
	http.HandleFunc("/sample.png", blockPic)
	http.ListenAndServe(":8080", nil)
}

func block(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset = utf-8")
	io.WriteString(w, `
	<img src="/sample.png">
	`)
}

func blockPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("sample.png")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
