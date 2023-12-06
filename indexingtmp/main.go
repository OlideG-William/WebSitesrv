package main

import (
	"html/template"
	"net/http"
)

var tpl, _ = template.New("myTemplate").Funcs(template.FuncMap{
	"lastItem": func(s []string) string {
		lastIndex := len(s) - 1
		return s[lastIndex]
	},
}).ParseFiles("frindex32.html")

var g []string

func main() {
	g = []string{"milk", "eggs", "green beans", "cheese", "flour", "sugar", "broccoli"}
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "frindex32.html", g)
}
