function initMap(){
    
    let map = new google.maps.Map(document.getElementById("map"), {
        zoom : 16,
        center : gps = {lat : 47.205617056212475, lng : -1.539378717397206},
        mapTypeID: "terrain",
    })

    new google.maps.Marker({
        position: gps,
        map: map,
    })
    // const icons = {
    //     locations : {
    //         icon : "/images/concerts.png"
    //     }
    // }
    // const features = [
    //     {
    //         position : new google.maps.gps(34.003342, -118.485832),
    //         type : "info",
    //     }
    // ]
    // for (let i = 0; i < features.length; i++) {
    //     const marker = new google.maps.Marker({
    //     position: features[i].position,
    //     icon: icons[features[i].type].icon,
    //     map: map,
    //     });
    // }
}