package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var tpl, _ = template.New("index.html").Funcs(template.FuncMap{
	"CanCashPr": func(p float64) string {
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
	},
	"Upper": strings.ToUpper,
}).ParseFiles("index.html")

var p float64

func main() {
	fmt.Println("tpl.Tree", *tpl.Tree)
	p = 3.33
	tpl, _ = tpl.ParseFiles("index.html")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8090", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", p)
}
