function initMap(){
    navigator.geolocation.getCurrentPosition(pos => {
        
        let crd = pos.coords
        
        lat = crd.latitude;
        lng = crd.longitude;
        
        console.log("Latitude : " + lat)
        console.log("Longitude : " + lng)
        
        let map = new google.maps.Map(document.getElementById("map"), {
            zoom : 10,
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
    
    address.addEventListener("change", function(event){
        
       
        let addressValue = event.target.value
        console.log(addressValue)
        
        // let position = address.geometry.location
        // console.log(position)
        
        // const address = document.getElementById("address").value;
        console.log(geocoder)
        geocoder.geocode({ address: addressValue }, (results, status) => {
            
            if (status === "OK") {
                resultsMap.setCenter(results[0].geometry.location);
                
                const image = "https://img.icons8.com/ios/452/concert.png"
                console.log(image)
                marker.setPosition(results[0].geometry.location),
                marker.icon = image
                
            } else {
                alert("Error: " + status);
            }
            // marker.addEventListener("click", toggleBounce)
        });
    });
}
