package data
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Group struct {
	ID int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Members []string `json:"members"`
	CreationDate int `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`
}

func GetGroups()[]Group {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	fmt.Print("Lancement Data GROUPS OK - ✅\n")

	spaceClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 2 seconds
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

	allGroups := []Group{}
	jsonErr := json.Unmarshal(body, &allGroups)
	// json.marshal à écrire dans le responsewriter
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	// fmt.Print(allGroups)
	
	return allGroups
}