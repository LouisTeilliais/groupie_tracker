package handlers

import (
	"fmt"
	"text/template"
	"net/http"
)


func Accueil(w http.ResponseWriter, req *http.Request){
	t := template.Must(template.ParseFiles("./template/home.html", "./template/layout/header.html", "./template/layout/footer.html"))

	fmt.Print("Accueil - ✅\n")
	if req.URL.Path == "/" {
        t.Execute(w, nil)
    } else if req.URL.Path != "/" {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }
	// templ , err := template.ParseFiles("./template/home.html")
	// if err != nil {
	// 	http.Error(w, "Internal Serveur Error", http.StatusInternalServerError)
	// 	return
	// }
	// templ.Execute(w, )
}