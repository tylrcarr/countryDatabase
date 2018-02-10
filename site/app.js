var map, geojson;
$(document).ready(function(){
	var osmBase = L.tileLayer('http://{s}.tile.osm.org/{z}/{x}/{y}.png');

	var baseMaps = {
		"OSMBaseLayer": osmBase
	};
	map = L.map('map', {
		layers: [osmBase],
		zoom: 5
	});
});
$.ajaxSetup({cache: false});
function getCountry()
{
	$.get("/country/" + document.getElementById("country").value, function(res) {
		console.log(res);
		$("#countryData").append("<img width='50%' src='" + res.flag + "' />");
		//$("#countryData").append("<div id='map'></div>");
		map.addLayer(L.geoJSON(res.geojson));
		map.setView({lat: res.latlng[0], lng: res.latlng[1]});
		$("#map").height($("#data").height() * .5);
		$("#outer").fadeOut();
		$("#data").fadeIn();
		map.invalidateSize();
		$("#backBtn").css('height', $("#backBtn").width());
	});
}

$("#data").resize(function (){
	$("#map").height($(this).height() * .5);
	map.invalidateSize();
	$("#backBtn").css('height', $("#backBtn").width());
});
