package data 

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
	"time"
)

type CityStruct struct {
	City []string  `json:"locations"`
}

func GetCity(thisID string)[]string{

	url := "https://groupietrackers.herokuapp.com/api/locations/" + thisID
	// fmt.Print(url)
	
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	fmt.Print(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	
	req.Header.Set("Page-Locations", "All-City")
	
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
	// fmt.Print(string(body))

	thisCity := CityStruct{}
	jsonErr := json.Unmarshal(body, &thisCity)
	
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return thisCity.City
}
