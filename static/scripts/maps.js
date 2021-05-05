function initMap(){
    navigator.geolocation.getCurrentPosition(pos => {
        
        let crd = pos.coords
        
        lat = crd.latitude;
        lng = crd.longitude;
        
        console.log("Latitude : " + lat)
        console.log("Longitude : " + lng)
        
        let map = new google.maps.Map(document.getElementById("map"), {
            zoom : 7,
            center : gps = {lat, lng},
        });
        
        const marker = new google.maps.Marker({
            position: gps,
            map: map,
        });
        
        const geocoder = new google.maps.Geocoder();
        codeAddress(geocoder, map, marker);
        
    })
    
}

function codeAddress(geocoder, resultsMap, marker) {
    
    const address = document.querySelector('select')
    
    address.addEventListener("change", async function(event){ 
        
        let addressValue = event.target.value
        // console.log(addressValue)
        
        let res = await fetch(`/cities?groups=${addressValue}`)
        let cities = await res.json()
        console.log(cities)
        
        const divCities = document.getElementById("villes_artistes")
        while (divCities.firstChild){
            divCities.removeChild(divCities.firstChild);
        }
        // divCities.append("Les concerts de votre artiste :")
        for (let i = 0; i < cities.length; i++){
            
            let replaceUnderscore = cities[i].replaceAll("_", " ")
            let replace = replaceUnderscore.replace("-", " ")
            const capitalize = replace.replace(/(^\w{1})|(\s+\w{1})/g, letter => letter.toUpperCase());
            let city = document.createElement("p");
            city.textContent = capitalize
            divCities.append(city)
        }
        console.log(divCities)
        
        for (let i = 0; i < cities.length; i ++) {
            console.log(cities.length)
            
            geocoder.geocode({ address: cities[i]}, (results, status) => {
                
                if (status === "OK") {
                    
                    // console.log(results[i])
                    resultsMap.setCenter(results[0].geometry.location);
                    
                    marker.setPosition(results[0].geometry.location)
                    
                } else {
                    alert("Error: " + status);
                }
            });
        }
    });
}
