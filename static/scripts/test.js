var yahoo = [], j = `[{"name":"Lenovo Thinkpad 41A4298","website":"google"},
{"name":"Lenovo Thinkpad 41A2222","website":"google"},
{"name":"Lenovo Thinkpad 41Awww33","website":"yahoo"},
{"name":"Lenovo Thinkpad 41A424448","website":"google"},
{"name":"Lenovo Thinkpad 41A429rr8","website":"ebay"},
{"name":"Lenovo Thinkpad 41A429ff8","website":"ebay"},
{"name":"Lenovo Thinkpad 41A429ss8","website":"rediff"},
{"name":"Lenovo Thinkpad 41A429sg8","website":"yahoo"}]`

var data = JSON.parse(j, function(key, value) { 
      if ( value.website === "yahoo" ) yahoo.push(value); 
      return value; })

console.log( yahoo )