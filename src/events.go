package handlers

import (
	"fmt"
	"text/template"
	"net/http"
)

func Events(w http.ResponseWriter, req *http.Request){
	t := template.Must(template.ParseFiles("./template/events.html", "./template/layout/header.html"))
	fmt.Print("Evenements - ✅\n")
    t.Execute(w, nil)
}