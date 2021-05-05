package handlers

import (
	"fmt"
	"text/template"
	"net/http"
)

func Locations(w http.ResponseWriter, req *http.Request){
	// t := template.Must(template.ParseFiles())

	const path = "./template/locations.html"

	t := template.Must(template.ParseFiles(path, "./template/layout/header.html", "./template/layout/footer.html"))

	fmt.Print("Locations - ✅\n")

	//gestion de l'erreur 500
	templ , err := template.ParseFiles(path)	//verification de la validiter de la page html
	if err != nil {
		http.Error(w, "Internal Serveur Error 500", http.StatusInternalServerError)
		return
	}else{
		t.Execute(w, nil)
		templ = templ
	}

}