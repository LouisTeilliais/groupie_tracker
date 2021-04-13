package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var t *template.Template
var t2 *template.Template
var t3 *template.Template
var t4 *template.Template

func accueil(w http.ResponseWriter, req *http.Request) {
	fmt.Print("accueil")

	// if req.Method == "GET" {
	t.Execute(w, nil)

	// 	if req.URL.Path != "/" {
	// 		http.Error(w, "404 not found.", http.StatusNotFound)
	// 	}
	// }

}

func artistes(w http.ResponseWriter, req *http.Request) {
	fmt.Print("artistes")
	t2.Execute(w, nil)
}

func villes(w http.ResponseWriter, req *http.Request) {
	fmt.Print("villes")
	t3.Execute(w, nil)
}

func concerts(w http.ResponseWriter, req *http.Request) {
	fmt.Print("concerts")
	t4.Execute(w, nil)
}

func main() {
	t = template.Must(template.ParseFiles("./template/index.html"))
	t2 = template.Must(template.ParseFiles("./template/artistes.html"))
	t3 = template.Must(template.ParseFiles("./template/villes.html"))
	t4 = template.Must(template.ParseFiles("./template/concerts.html"))
	http.HandleFunc("/", accueil)
	http.HandleFunc("/artistes", artistes)
	http.HandleFunc("/villes", villes)
	http.HandleFunc("/concerts", concerts)
	http.ListenAndServe(":8000", nil)

}
