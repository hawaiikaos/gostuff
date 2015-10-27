console.log("in main.js");

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
        getCountry();
    };
    
    function error() {
        output.innerHTML = "Unable to locate you";
    };
    
    output.innerHTML = "Locating...";
    
    navigator.geolocation.getCurrentPosition(success, error);
}