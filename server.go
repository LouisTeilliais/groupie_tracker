package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var t *template.Template
var t2 *template.Template

func accueil(w http.ResponseWriter, req *http.Request) {
	fmt.Print("hello")
	t.Execute(w, nil)
}

func artistes(w http.ResponseWriter, req *http.Request) {
	fmt.Print("hello2")
	t2.Execute(w, nil)
}

func main() {
	t = template.Must(template.ParseFiles("./template/index.html"))
	t2 = template.Must(template.ParseFiles("./template/artistes.html"))
	http.HandleFunc("/", accueil)
	http.HandleFunc("/artistes", artistes)
	http.ListenAndServe(":8000", nil)

}
