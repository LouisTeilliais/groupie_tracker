const artistes = "https://cors-anywhere.herokuapp.com/https://groupietrackers.herokuapp.com/api/artists"
const locations = "https://groupietrackers.herokuapp.com/api/locations"
const dates = "https://groupietrackers.herokuapp.com/api/dates"
const relations = "https://groupietrackers.herokuapp.com/api/relation"
function dataBase(test){
    console.log(test[0].members[0])
    console.log(test[0].members[1])
    console.log(test[0].members[2])
}

fetch(api)
.then((response) => response.json())
.then(dataBase)
