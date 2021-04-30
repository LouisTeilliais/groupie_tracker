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

	t := template.Must(template.ParseFiles("./template/artistes.html", "./template/layout/header.html"))
	fmt.Print("Artistes - ✅\n")
	// w.Header().Add("Content-Type", "application/json")
    t.Execute(w, pageGroups)
}


