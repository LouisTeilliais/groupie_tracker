package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type group struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Members []string `json:"members"`
	CreationDate int `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`
}

func main(){

	url := "https://groupietrackers.herokuapp.com/api/artists"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Page-Artists", "All-Groups")

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

	allGroups := []group{}
	var groupsList []group = []group{}
	jsonErr := json.Unmarshal(body, &allGroups)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	for i:=0; i<50; i++{
		groupsList = append(groupsList, allGroups[i])
		fmt.Printf("ID : %v - ", allGroups[i].ID)
		fmt.Printf("Nom : %v - ", allGroups[i].Name)
		fmt.Printf("Création : %v - ", allGroups[i].CreationDate)
		fmt.Printf("Premier Membre : %v\n", allGroups[i].Members[0])
	}

	fmt.Printf("test groupe : %v\n", groupsList[0])
}