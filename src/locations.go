package handlers

import (
	"fmt"
	"text/template"
	"encoding/json"
	"net/http"
	data "./data"
)

func Locations(w http.ResponseWriter, req *http.Request){
	
	groups := data.GetGroups()
	t := template.Must(template.ParseFiles("./template/locations.html", "./template/layout/header.html", "./template/layout/footer.html"))
	
	if req.Method == "POST" {
		
		fmt.Print("req OK", "\n")
		// indexGroup := req.FormValue("groups")
		// fmt.Print("coucou")
		// cities := data.GetCity(indexGroup)
		// fmt.Print(cities)
	}
		// idCity := req.FormValue("city")
		
		type DataCity struct {
			// Cities []string
			// ThisGroup data.OneGroup
			Groups []data.Group
		}
		
		pageLocations := DataCity{Groups: groups}
		t.Execute(w, pageLocations)
		
		fmt.Print("Locations - ✅\n")
		
	
}

func ConcertsLocations(w http.ResponseWriter, req *http.Request){

	// groups := data.GetGroups()
	indexGroup := req.FormValue("groups")
	cities := data.GetCity(indexGroup)

	jsonData, _ := json.Marshal(cities)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}