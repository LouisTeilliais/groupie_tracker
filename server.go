package main

import (
<<<<<<< HEAD
	"net/http"
	handlers "./src"
)

=======
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template
var t2 *template.Template
var t3 *template.Template
var t4 *template.Template

func accueil(w http.ResponseWriter, req *http.Request) {
	fmt.Print("accueil")
	if req.URL.Path == "/" {
		t.Execute(w, nil)
	} else if req.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
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

>>>>>>> Louis
func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.Accueil)
	http.HandleFunc("/artistes", handlers.Artistes)
	http.HandleFunc("/locations", handlers.Locations)
	http.HandleFunc("/events", handlers.Events)
	http.ListenAndServe(":8000", nil)

}
