const locations = "https://cors-anywhere.herokuapp.com/https://groupietrackers.herokuapp.com/api/locations"
// const relations = "https://cors-anywhere.herokuapp.com/https://groupietrackers.herokuapp.com/api/relation"

function Locations(locations){
    console.log(locations.index)
}
fetch(locations)
.then((response) => response.json())
.then(Locations)