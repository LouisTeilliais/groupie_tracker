package handlers

import (	//Importation des librairies nécessaire
	"fmt"
	"text/template"
	"net/http"
)


func Accueil(w http.ResponseWriter, req *http.Request){		//fonction de qui affiche la page Accueil

	//importation des page HTML 

	t := template.Must(template.ParseFiles("./template/home.html", "./template/layout/header.html", "./template/layout/footer.html"))

	fmt.Print("Accueil - ✅\n")

	//gestion de l'erreur 404

	if req.URL.Path == "/" {	//verification de l'URL
        t.Execute(w, nil)
    } else if req.URL.Path != "/" {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }
	

	//gestion de l'erreur 500

	templ , err := template.ParseFiles("./template/home.html")	//verification de la validiter de la page html
	if err != nil {
		http.Error(w, "Internal Serveur Error 500", http.StatusInternalServerError)
		return
	}else{
		templ.Execute(w, nil)
	}
	
}