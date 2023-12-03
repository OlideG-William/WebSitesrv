package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type prodSpec struct {
	Size   string
	Weight float32
	Descr  string
}

type product struct {
	ProdID int
	Name   string
	Cost   float64
	Specs  prodSpec
}

var tpl *template.Template
var prod1 product

func main() {
	prod1 = product{
		ProdID: 5,
		Name:   "Wicked Coll Phone",
		Cost:   899,
		Specs: prodSpec{
			Size:   "150 x 70 x 7 mm",
			Weight: 65,
			Descr:  "Over priced shiny thing designed to shatter on impact ",
		},
	}

	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/productinfo2", productInfoHandler)
	http.ListenAndServe(":8080", nil)

}

func productInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("indeHandler running")

	tpl.ExecuteTemplate(w, "productinfo2.html", prod1)
}
