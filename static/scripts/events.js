console.log("test")
allGroups= document.querySelectorAll(".groups")
allGroups.forEach(function(oneGroup){
    oneGroup.addEventListener('click', function(){
    // let lastGroup = document.querySelector('.clické');
    // lastGroup.classList.remove('clické');
    oneGroup.setAttribute('class', 'clické');
    })
})


// let form = document.getElementById("form");
// let test = Array.from(document.querySelectorAll('.groups'))
// test.forEach(function(thisGroup){
//     thisGroup.addEventListener('click', function(){
//     getDates(thisGroup.getAttribute("id"))
//     })
// })

// form.addEventListener("submit", function (event) {
//     event.preventDefault();
//     const searchSubmitted = event.target[0].value
//     searchGroups(searchSubmitted)
// })

// function searchGroups(text){
//     let regGroup = new RegExp(text, 'i')
//     test.forEach(function(test){
//         let eachGroup = test.innerHTML
//         if(regGroup.test(eachGroup)){
//             test.setAttribute('class', "groups")
//         } else {
//             test.setAttribute('class', "groups hidden")
//             console.log("pas compatible")
//         }
//     })
// }


// let dates = document.querySelectorAll('.dates')

// function getDates(idGroup){
//     dates.forEach(function(thisDate){
//         if (thisDate.getAttribute("id")==idGroup){
//             let thisDateStr = thisDate.innerHTML
//             let tabDates = thisDateStr.substr(1, thisDateStr.length-2).split(' ')
//             console.log(tabDates)
//             thisDate.setAttribute("class", "dates")
//         } else {
//             thisDate.setAttribute("class", "dates hidden")
//         }
//     })
// }