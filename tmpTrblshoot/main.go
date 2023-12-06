package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type price struct {
	UppercaseField float64
	lowercaseField float64
}

var FuncMap = template.FuncMap{
	"CanCashPr": CanCashPr,
	"Upper":     strings.ToUpper,
}
var tpl *template.Template
var p price

func main() {

	tpl = template.New("index.html")

	tpl = tpl.Funcs(FuncMap)

	tpl = template.Must(tpl.ParseFiles("templates/index.html"))

	p.UppercaseField = 3.3333
	p.lowercaseField = 3.3333

	http.HandleFunc("/index1", indexHandler1)
	http.HandleFunc("/index2", indexHandler2)
	http.HandleFunc("/error", errorHandler)
	http.ListenAndServe("localhost:8090", nil)
}

func CanCashPr(p float64) (str string, err error) {
	err = errors.New("warning")
	remainder := int(p*100) % 5
	quotiant := int(p*100) / 5
	if remainder < 3 {
		pr := float64(quotiant*5) / 100
		s := fmt.Sprintf("%.2f", pr)
		return s, err
	}
	pr := (float64(quotiant*5) + 5) / 100
	s := fmt.Sprintf("%.2f", pr)
	return s, err
}

func indexHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-------------indexHandler1 running-----------")

	tpl.ExecuteTemplate(w, "index.html", p)
}

func indexHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-------------indexHandler2 running-----------")
	var buff bytes.Buffer

	err := tpl.Execute(&buff, p)
	if err != nil {
		fmt.Println("err:", err)
		http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
		return
	}
	buff.WriteTo(w)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "there was an error parsing page")
}
