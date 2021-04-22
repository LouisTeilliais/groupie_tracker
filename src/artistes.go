package handlers

import (
	"text/template"
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

type PageDataGroups struct {
	Groups []group
}

func Artistes(w http.ResponseWriter, req *http.Request){
	groups := GetGroups()
	pageGroups := PageDataGroups{Groups: groups}
	fmt.Printf("test groupe : %v\n", pageGroups.Groups[0])

	t := template.Must(template.ParseFiles("./template/artistes.html", "./template/layout/header.html"))
	// fmt.Print("Artistes - ✅\n")
	// w.Header().Add("Content-Type", "application/json")
    t.Execute(w, pageGroups)
}



func GetGroups()[]group {
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
	jsonErr := json.Unmarshal(body, &allGroups)
	// json.marshal à écrire dans le responsewriter
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Print(allGroups)

	return allGroups
}