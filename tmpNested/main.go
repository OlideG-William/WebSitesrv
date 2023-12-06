package main

import (
	"html/template"
	"net/http"
)

type task struct {
	Name string
	Done bool
}

var tpl *template.Template

func main() {
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8090", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index2.html", nil)
}
