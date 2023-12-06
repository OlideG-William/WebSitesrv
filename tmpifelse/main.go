package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name     string
	Language string
	Member   bool
}

var u User
var tpl *template.Template

func main() {
	//u = User{"Bob Smith", "English", false}
	// u = User{"Juan Hernandez", "Spanish", true}
	u = User{"Zhan Jang", "Mandarin", true}
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/welcome", welcomeHandler)
	http.ListenAndServe(":8090", nil)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "membership2.html", u)
}
