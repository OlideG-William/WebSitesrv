package main

import (
	"html/template"
	"net/http"
)

type GroceryList []string

var tpl *template.Template
var g GroceryList

func main() {
	g = GroceryList{"milk", "eggs", "green beans", "cheese", "flour", "sugar", "broccoli"}
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/list", listHandler)
	http.ListenAndServe(":8090", nil)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "wrong.html", g)
}
