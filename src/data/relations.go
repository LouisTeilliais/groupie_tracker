package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"strings"
)
type Infos struct {
	ID int `json:"id"`;
	TabEvents interface{} `json:"datesLocations"`;
}
type Relations struct {
	City string;
	Country string;
	Dates []string;
}
	
func GetEvents(thisID string)[]Relations{
	url := "https://groupietrackers.herokuapp.com/api/relation/"+ thisID
	fmt.Print("Data Relations OK - ✅\n")
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	
	req.Header.Set("Page-Events", "All-Dates")
	
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	
	if res.Body != nil {
		defer res.Body.Close()
	}
	
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	
	infosTab := Infos{}
	
	jsonErr := json.Unmarshal(body, &infosTab)
	// json.marshal à écrire dans le responsewriter
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	mapInfos := infosTab.TabEvents.(map[string]interface{})
	var tabRelations []Relations
	// création d'un tableau de structures avec les villes et les dates.
	//on y met la clé et la valeur.
	for k, v := range mapInfos{
		tabLocation := strings.Split(k, "-")
		city := strings.Title(strings.Replace(tabLocation[0], "_", " ", -1))
		country := strings.Title(strings.Replace(tabLocation[1], "_", " ", -1))


		var tabValues []string
		 for _, v2 := range v.([]interface{}){
			tabValues = append(tabValues, v2.(string))
		 }
		// tableau de Relations dans lequel je push ça)
		var oneRelation = Relations{city, country, tabValues}
		tabRelations = append(tabRelations, oneRelation)
	}
// strings.replace("telle chaine", " espace ")
// bibliothèque strings
// Go capitalize = strings.title majuscule à chaque nouveau mot
// cournty city dates et k.split(k, " ") pour séparer les 2 ! prmeier country, second ville
	return tabRelations
}