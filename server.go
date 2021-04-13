package main

import (
	"net/http"

	handlers "./src"
)

func main() {
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.Accueil)
	// http.HandleFunc("/artistes", artistes)
	// http.HandleFunc("/villes", villes)
	// http.HandleFunc("/concerts", concerts)
	http.ListenAndServe(":8000", nil)

}
