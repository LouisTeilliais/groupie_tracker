//Sélection de la barre de recherche sur la page html
const form = document.getElementById('formSearch') 

// Redirection vers le groupe cliqué
function displayGroup(ID){
    // Ajout d'une action au form pour la redirection vers l'ID du groupe
    form.setAttribute('action', `/events?oneGroup=${ID}`)
    // Submit du form
    form.submit()
}