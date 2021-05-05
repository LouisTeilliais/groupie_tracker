package handlers

import (
	"text/template"
	"fmt"
	"net/http"
	data "./data"
)

type PageDataGroups struct {
	Groups []data.Group
}

func Artistes(w http.ResponseWriter, req *http.Request){
	groups := data.GetGroups()
	pageGroups := PageDataGroups{Groups: groups}
	
	const path = "./template/artistes.html"

	t := template.Must(template.ParseFiles(path, "./template/layout/header.html"))
	fmt.Print("Artistes - ✅\n")
	

	//gestion de l'erreur 500
	templ , err := template.ParseFiles(path)	//verification de la validiter de la page html
	if err != nil {
		http.Error(w, "Internal Serveur Error 500", http.StatusInternalServerError)
		return
	}else{
		t.Execute(w, pageGroups)
		templ = templ
	}
}