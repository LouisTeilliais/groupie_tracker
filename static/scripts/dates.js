
const dates = "https://cors-anywhere.herokuapp.com/https://groupietrackers.herokuapp.com/api/dates"
const relations = "https://cors-anywhere.herokuapp.com/https://groupietrackers.herokuapp.com/api/relation"
function Dates(dates){
    console.log(dates.index)
}
fetch(relations)
.then((response) => response.json())
.then(Dates)