package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main(){
	GetEvents("1")
}
type infos struct {
	ID int `json:"id"`;
	TabEvents interface{} `json:"datesLocations"`;
}
	
func GetEvents(thisID string)interface{} {
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
	
	tabtest := infos{}
	// tabIndex2 := index2{}
	// fmt.Print(tabIndex2)
	
	jsonErr := json.Unmarshal(body, &tabtest)
	// json.marshal à écrire dans le responsewriter
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Printf("tabtest : %v\n", tabtest)
	fmt.Printf("tabtest.TabEvents : %v\n", tabtest.TabEvents)

	fmt.Print("Fin GetEvents OK - ✅\n")
	return tabtest.TabEvents
}