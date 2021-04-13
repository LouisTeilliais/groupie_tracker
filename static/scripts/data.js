
function dataBase(test){
    console.log(test)
}


fetch('https://groupietrackers.herokuapp.com')
.then((response) => response.json())
.then(dataBase)

