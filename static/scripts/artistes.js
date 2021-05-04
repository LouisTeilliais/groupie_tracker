let start = document.getElementById("startCD") 
let from = document.getElementById("fromCD")
from.innerHTML = start.value

start.addEventListener('input', function func(e){
    from.innerHTML = e.target.value
})

let end = document.getElementById("endCD") 
let to = document.getElementById("toCD")
to.innerHTML = end.value

end.addEventListener('input', function func(e){
    to.innerHTML = e.target.value
})

// FIRST ALBUM

let start2 = document.getElementById("startFA") 
let from2 = document.getElementById("fromFA")
from2.innerHTML = start2.value

start2.addEventListener('input', function func(e){
    from2.innerHTML = e.target.value
})

let end2 = document.getElementById("endFA") 
let to2 = document.getElementById("toFA")
to2.innerHTML = end2.value

end2.addEventListener('input', function func(e){
    to2.innerHTML = e.target.value
})