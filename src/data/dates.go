package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Date struct {
	ID int `json:"id"`
	TabDates []string `json:"dates"`
}

// type index struct {
// 	Index []Date `json:"Index"`
// }

func GetDates(thisID string)Date {
	url := "https://groupietrackers.herokuapp.com/api/dates/" +thisID
	fmt.Print("Lancement Data DATES OK - ✅\n")
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
	
	tabIndex := Date{}
	jsonErr := json.Unmarshal(body, &tabIndex)
	// json.marshal à écrire dans le responsewriter
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	
	return tabIndex
}