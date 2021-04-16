package handlers

import (
	"fmt"
	"text/template"
	"net/http"
)

func Artistes(w http.ResponseWriter, req *http.Request){
	t := template.Must(template.ParseFiles("./template/artistes.html", "./template/layout/header.html"))
	fmt.Print("Artistes - ✅\n")
    t.Execute(w, nil)
}