package handlers

import (
	"fmt"
	"text/template"
	"net/http"
)

func Locations(w http.ResponseWriter, req *http.Request){
	// t := template.Must(template.ParseFiles())

	t := template.Must(template.ParseFiles("./template/locations.html", "./template/layout/header.html", "./template/layout/footer.html"))

	fmt.Print("Locations - ✅\n")
        t.Execute(w, nil)

}