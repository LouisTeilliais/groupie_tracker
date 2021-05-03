package handlers

import (
	"fmt"
	"text/template"
	"net/http"
	"strconv"
	data "./data"
)

func Locations(w http.ResponseWriter, req *http.Request){
	
	
	t := template.Must(template.ParseFiles("./template/locations.html", "./template/layout/header.html", "./template/layout/footer.html"))
	
	var AllCities []string
	var cities []string

	for indexGroup := 1; indexGroup < 52; indexGroup++ {
		
		cities = data.GetCity(strconv.Itoa(indexGroup))

		for _, value := range cities{
			
			for j := 0; j < len(cities); j++ {
				
				if value != cities[j]{
					AllCities = append(AllCities, value)
				}
			}	
			fmt.Print(AllCities)		
		}	
	}

	if req.Method == "POST" {
		// idCity := req.FormValue("city")
		
		type DataCity struct {
			Cities []string
		}
		
		pageLocations := DataCity{Cities : cities}
		t.Execute(w, pageLocations)
	}else {
		
		fmt.Print("Locations - ✅\n")
		t.Execute(w, nil)
	}
}