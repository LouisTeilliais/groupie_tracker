package data
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type OneGroup struct {
	ID int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Members []string `json:"members"`
	CreationDate int `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`
}

func GetOneGroup(thisID string)OneGroup {
	url := "https://groupietrackers.herokuapp.com/api/artists/"+thisID
	fmt.Print("Lancement Data ONE GROUP OK - ✅\n")
	
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
	
	thisGroup := OneGroup{}
	jsonErr := json.Unmarshal(body, &thisGroup)
	// json.marshal à écrire dans le responsewriter
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	// fmt.Print(allGroups)
	
	return thisGroup
}