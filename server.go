package main

//importation des librairies utilisées
import (
	"net/http"
	"fmt"
	handlers "./src"
)


func main() {
	fmt.Print("Démarrage du serveur... 💬\n")

	// Récupération des fichiers static pour l'affichage des pages
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	//référencement des pages html
	http.HandleFunc("/", handlers.Accueil)
	http.HandleFunc("/artistes", handlers.Artistes)
	http.HandleFunc("/locations", handlers.Locations)
	http.HandleFunc("/events", handlers.Events)
	http.HandleFunc("/cities", handlers.ConcertsLocations)
	
	//assignation du port du serveur
	http.ListenAndServe(":8000", nil)
	
}

