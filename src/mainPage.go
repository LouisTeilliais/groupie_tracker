package handlers

import (
	"fmt"
	"text/template"
	"net/http"
)
func Accueil(w http.ResponseWriter, req *http.Request){
	t := template.Must(template.ParseFiles("./template/index.html"))
	fmt.Print("accueil")
	if req.URL.Path == "/" {
        t.Execute(w, nil)
    } else if req.URL.Path != "/" {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }
	fmt.Print("test")
}