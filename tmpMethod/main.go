package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template
var p Price

type Price float64

func (p Price) CanCashPr() string {
	remainder := int(p*100) % 5
	quotiant := int(p*100) / 5
	if remainder < 3 {
		pr := float64(quotiant*5) / 100
		s := fmt.Sprintf("%.2f", pr)
		return s
	}
	pr := (float64(quotiant*5) + 5) / 100
	s := fmt.Sprintf("%.2f", pr)
	return s
}

func main() {
	p = 1.4
	tpl, _ = template.ParseFiles("index.html")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":9080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", p)
}
