package handlers

import (
	"text/template"
	"fmt"
	"net/http"
	data "./data"
)
func Events(w http.ResponseWriter, req *http.Request){
	groups := data.GetGroups()
	t := template.Must(template.ParseFiles("./template/events.html", "./template/layout/header.html"))
	
	if req.Method == "POST" {
		idGroup := req.FormValue("groups")
		fmt.Printf("Requête POST OK | ID du groupe : %s | OK - ✅\n", idGroup)
		

		// fmt.Printf("Retour GetEvents : %s", relations)
		type PageDataDates struct {
			Dates data.Date
			ThisGroup data.OneGroup
			Groups []data.Group
			Relations []data.Relations
		}

		relations := data.GetEvents(idGroup)
		dates := data.GetDates(idGroup)
		group := data.GetOneGroup(idGroup)
		// fmt.Printf("dates %v", dates)
		pageDates := PageDataDates{Dates: dates, Groups: groups, ThisGroup: group, Relations:relations}
		t.Execute(w, pageDates)
		//w.Headers().Add("Content-Type", "application/json") ajouter des headers
		//w.Write chercher doc, ajouter les données
	} else {
		type PageDataDates struct {
			Groups []data.Group
		}
		pageDates2 := PageDataDates{Groups: groups}
		fmt.Print("Events - ✅\n")
		t.Execute(w, pageDates2)
	}
}
