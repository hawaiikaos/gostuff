console.log("in main.js");

//var boundsCollection = new Object();
boundsCollection = [];
 
function reverseGeocode(lat,long) {
    address = "United States";
    //$.getJSON( "https://maps.googleapis.com/maps/api/geocode/json?latlng="+lat+","+long+"&key=AIzaSyAhPI53u5k-IJ80w1ISj_7QZdxKhQiCpTw",
    $.getJSON( "https://maps.googleapis.com/maps/api/geocode/json?address="+address+"&key=AIzaSyAhPI53u5k-IJ80w1ISj_7QZdxKhQiCpTw",
     function( data ) {
         $.each( data, function( key, val ) {
             console.log(key, ": ", val);
             /*for (i = 0; i < val.length; i++){
                 console.log(val[i].address_components);
                 for (j = 0; j < val[i].address_components.length; j++) {
                     console.log(val[i].address_components[j].types[0]);
                     console.log(val[i].address_components[j].long_name);
                 }
             }*/
         });
     });
}

function boundViewer(entity) {
    var bounds = new Object();
    $.getJSON( "https://maps.googleapis.com/maps/api/geocode/json?address="+entity+"&key=AIzaSyAhPI53u5k-IJ80w1ISj_7QZdxKhQiCpTw",
     function( data ) {
         $.each( data, function( key, val ) {
             console.log(key, ": ", val);
             try {
                 console.log(val[0].geometry.bounds);
                 //bounds.push(entity);
                 //bounds.push(val[0].geometry.bounds);
                 bounds.bounds = val[0].geometry.bounds;
                 bounds.region_name = entity;
             } catch (e) {
                 //ignore
             }
             
             /*for (i = 0; i < val.length; i++){
                 console.log(val[i].address_components);
                 for (j = 0; j < val[i].address_components.length; j++) {
                     console.log(val[i].address_components[j].types[0]);
                     console.log(val[i].address_components[j].long_name);
                 }
             }*/
         });
     });
     return bounds;
}

function findHemisphere(lat,long) {
    var hemisphere = [2];
    if (lat > 0) {
        hemisphere[0] = "Northern";
    } else {
        hemisphere[0] = "Southern";
    }
    if (long > 0) {
        hemisphere[1] = "Eastern";
    } else {
        hemisphere[1] = "Western";
    }
    return(hemisphere);
}

function compareLatNorth(lat1, lat2) {
    if (lat1 > lat2) {
        return lat1;
    } else {
        return lat2;
    }
}

function compareLatSouth(lat1, lat2) {
    if (lat1 < lat2) {
        return lat1;
    } else {
        return lat2;
    }
}

function compareLongEast(long1, long2) {
    if (long1 > long2) {
        return long1;
    } else {
        return long2;
    }
}

function compareLongWest(long1, long2) {
    if (long1 < long2) {
        return long1;
    } else {
        return long2;
    }
}

function getCountry() {
    $.getJSON( "/json/world-countries.json", function( data ) {
      var items = [];
      $.each( data, function( key, val ) {
          console.log(val.length);
          try {
              //console.log(val[0].properties.name);
              //console.log(val[0].geometry.coordinates);
              for (i = 0; i < val.length; i++) {
                  console.log(val[i].properties.name);
                  //console.log(val[i].geometry.coordinates.length);
                  for (j = 0; j < val[i].geometry.coordinates.length; j++) {
                      //console.log(val[i].geometry.coordinates[j]);
                      for (k = 0; k < val[i].geometry.coordinates[j].length; k++) {
                          //console.log(val[i].geometry.coordinates[j][k][0]);
                          console.log(val[i].geometry.coordinates[j][k].length);
                          for (m = 0; m < val[i].geometry.coordinates[j][k].length; m++) {
                              //console.log(val[i].geometry.coordinates[j][k][m]);
                              var coord = val[i].geometry.coordinates[j][k][m];
                              var oldcoord = val[i].geometry.coordinates[j][k][m];
                              var hemisphere = findHemisphere(coord[1], coord[0]);
                              console.log(coord);
                              console.log(hemisphere[0], " ", hemisphere[1]);
                              if (m === 0) {
                                  console.log("first coordinate in set");
                              } else {
                                  oldcoord = val[i].geometry.coordinates[j][k][(m - 1)];
                              }
                          }
                      }
                  }
              }
          } catch(e){
              //do nothing
          }
      });
    });
}

function geoFind() {
    var output = document.getElementById("output");
    
    if (!navigator.geolocation) {
        output.innerHTML = "Geolocation is not supported by your browser";
        return;
    }
    
    function success(position) {
        var latitude = position.coords.latitude;
        var longitude = position.coords.longitude;
        
        var hemisphere = findHemisphere(position.coords.latitude, position.coords.longitude);
        output.innerHTML = 'Lat: ' + latitude + '<br>' + 'Lon: ' + longitude + '<br>' + hemisphere[0] + '<br>' + hemisphere[1];
        //getCountry();
        //reverseGeocode(latitude,longitude);
        //reverseGeocode(51.52,-100.42);
    };
    
    function error() {
        output.innerHTML = "Unable to locate you";
    };
    
    output.innerHTML = "Locating...";
    
    navigator.geolocation.getCurrentPosition(success, error);
    
    boundsCollection.push(boundViewer("England"));
    boundsCollection.push(boundViewer("Wales"));
}

function downloadBounds() {
    console.log(boundsCollection);
    //var csvContent = "";
    var datastring = "data:text/csv;charset=utf-8,";
    //boundsCollection.forEach(function(infoArray, index) {
        /*try {
            dataString = infoArray.join(",");
            csvContent += index < boundsCollection.length ? dataString+ "\n" : dataString;
        } catch (e) {
            //
        }*/
    //});
    
    for (i = 0; i < boundsCollection.length; i++) {
        console.log(boundsCollection[i]);
        datastring += boundsCollection[i].region_name;
        var itemBounds = boundsCollection[i].bounds;
        datastring += "," + itemBounds.northeast.lat + "," + itemBounds.northeast.lng;
        datastring += "," + itemBounds.southwest.lat + "," + itemBounds.southwest.lng + "\n";
    }
    var encodedUri = encodeURI(datastring);
    window.open(encodedUri);
}