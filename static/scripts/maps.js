// let geocoder;
let map;

function coordonnees(pos){
    let crd = pos.coords

    let lat = crd.latitude;
    let lng = crd.longitude;
    
    console.log("Latitude : " + lat)
    console.log("Longitude : " + lng)
    
    let gps = {lat, lng}
    
    let map = new google.maps.Map(document.getElementById("map"), {
        zoom : 10,
        center : gps,
        mapTypeID: "terrain",
    });

    new google.maps.Marker({
        position: gps,
        map: map,
    });
}
navigator.geolocation.getCurrentPosition(coordonnees)



function initMap(){
    
    
    // const geocoder = new google.maps.Geocoder();
    // document.getElementById("submit").addEventListener("click", () => {
    //   codeAddress(geocoder, map);
    // });
}
function codeAddress(geocoder, resultsMap) {
  
    const address = document.querySelector('select')
    
    address.addEventListener("change", function(event){
    
    // console.log(address)

    let addressValue = event.target.value
    console.log(addressValue)

//     geocoder.geocode( { 'address': addressValue}, (results, status) => {
//         console.log("test")
//         if (status == 'OK') {
//         resultsMap.setCenter(results[0].geometry.location);
//         var marker = new google.maps.Marker({
//             map: map,
//             position: results[0].geometry.location
//         });

//         marker.setMap(map)

//         }else {
//         alert('Error : ' + status);
//         }
//     });
  });
}
codeAddress()


