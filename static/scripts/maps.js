function LocationsMaps(){
    // let gps = {lat : 47.205617056212475, lng : -1.539378717397206}
    let map = new google.maps.Map(document.getElementById("map"), {
        zoom : 2,
        center : gps = {lat : 47.205617056212475, lng : -1.539378717397206},
        mapTypeID: "terrain",
    })

    let pointer = new google.maps.Marker({
        position: gps,
        map: map,
    })
}