$.ajaxSetup({cache: false});
var map, geojson, names;
$(document).ready(function(){
	var osmBase = L.tileLayer('http://{s}.tile.osm.org/{z}/{x}/{y}.png');

	var baseMaps = {
		"OSMBaseLayer": osmBase
	};
	map = L.map('map', {
		layers: [osmBase],
		zoom: 4
	});
	$.get("/names", function(res) {
		$("#country").autoComplete({
		    	minChars: 2,
			source: function(term, suggest){
				term = term.toLowerCase();
				var choices = res.names;
				var matches = [];
				for (i=0; i<choices.length; i++)
					if (~choices[i].toLowerCase().indexOf(term)) matches.push(choices[i]);
				suggest(matches);
			    	}
		});
	});
	$("#back").click(goBack);
});
function getCountry()
{
	$.get("/country/" + document.getElementById("country").value, function(res) {
		console.log(res);
		$("#countryData").append("<img id='flag' width='50%' src='" + res.flag + "' />");
		//$("#countryData").append("<div id='map'></div>");
		//geojson = L.geoJSON(res.geojson);
		$.get("/geojson/" + res.alpha3 + ".geo.json", function(res){
			console.log(res.features[0]);
			geojson = L.geoJSON(res.features[0]);	
			map.addLayer(geojson);
		});
		map.setView({lat: res.latlng[0], lng: res.latlng[1]});
		$("#map").height($("#data").height() * .5);
		$("#outer").fadeOut();
		$("#data").fadeIn();
		map.invalidateSize();
		$("#backBtn").css('height', $("#backBtn").width());
	});
}

function goBack() {
	$("#data").fadeOut();
	$("#outer").fadeIn();
	$("#flag").remove();	
	map.removeLayer(geojson);
}

$("#data").resize(function (){
	$("#map").height($(this).height() * .5);
	map.invalidateSize();
	$("#backBtn").css('height', $("#backBtn").width());
});
