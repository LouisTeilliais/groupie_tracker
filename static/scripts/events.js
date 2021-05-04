let form = document.getElementById('formSearch')

function displayGroup(ID){
    form.setAttribute('action', `/events?oneGroup=${ID}`)
    form.submit()
}