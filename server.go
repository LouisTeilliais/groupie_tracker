package main

import (
	"fmt"
	"text/template"
	"encoding/json"
	"strconv"
	"io/ioutil"
	"net/http"
	handlers "./src"
)

const RESTAPIURL string = "https://groupietrackers.herokuapp.com/api/artists"
// const RESTAPIURL string = "https://groupietrackers.herokuapp.com/api/artistes"

type Artist struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type PageDataArtist struct {
	Artists []Artist
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	http.HandleFunc("/", handlers.Accueil)
	http.HandleFunc("/artistes", handlers.Artistes)
	http.HandleFunc("/locations", handlers.Locations)
	http.HandleFunc("/events", handlers.Events)
	http.HandleFunc("/test", DefaultPage)
	
	http.ListenAndServe(":8000", nil)
	
}

func DefaultPage(w http.ResponseWriter, req *http.Request){
	
	var itemsPerPage int = 20
	var firstID int = 1
	
	artists := GetArtists(firstID, itemsPerPage)
	
	pageArtists := PageDataArtist{Artists: artists}
	
	t := template.Must(template.ParseFiles("./template/home.html", "./template/layout/header.html", "./template/layout/footer.html"))
	
	t.Execute(w, pageArtists)
	fmt.Print("TEST - ✅\n")
	fmt.Print(pageArtists)

	
	
}

func GetArtists(first int, nbItems int) []Artist{
	var artistList []Artist = []Artist{}
	
	i := first 
	for i < first+nbItems {
		res, err := http.Get(RESTAPIURL+"id/"+strconv.Itoa(i)+".json")
		if err != nil{
			fmt.Println(err)
			return nil
		}
		
		data, err := ioutil.ReadAll(res.Body)
		
		if err != nil{
			fmt.Println(err)
			return nil
		}
		
		var tmpArt Artist
		json.Unmarshal(data, &tmpArt)
		// if tmpArt.ID != 0 {
			fmt.Println(i)
			artistList= append(artistList, tmpArt)
			i++
		// }	
	}
	return artistList
}