package main

import (
	"net/http"
	"fmt"
	handlers "./src"
)


func main() {
	fmt.Print("Démarrage du serveur... 💬\n")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	http.HandleFunc("/", handlers.Accueil)
	http.HandleFunc("/artistes", handlers.Artistes)
	http.HandleFunc("/locations", handlers.Locations)
	http.HandleFunc("/events", handlers.Events)
	http.HandleFunc("/cities", handlers.ConcertsLocations)
	
	http.ListenAndServe(":8000", nil)
	
}

