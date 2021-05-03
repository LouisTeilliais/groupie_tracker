package main 

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CityStruct struct {
	City []string `json:"locations"`
}
func main(){
	GetCity()
}

func GetCity()CityStruct{
	url := "https://groupietrackers.herokuapp.com/api/locations/1"
	fmt.Print(url)
	
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	
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
	
	thisCity := CityStruct{}
	jsonErr := json.Unmarshal(body, &thisCity)
	
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	
	fmt.Print(thisCity)
	return thisCity
}
