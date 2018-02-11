var map, allGeojson, cSelect;
var cartoLight = L.tileLayer("https://cartodb-basemaps-{s}.global.ssl.fastly.net/light_all/{z}/{x}/{y}.png", {
  maxZoom: 19
});
L.Control.CountrySelector = L.Control.extend({
	vars: {
		div: ""
	},
	onAdd: function(map) {
		var div = L.DomUtil.create("div", "container");
		div.innerHTML = "<p class='h5'>Select Country</p><div class='radio'> <label><input id='sel1' onclick='changeCountry(1)' type='radio' name='optradio'> Country 1</label> </div> <div class='radio'> <label><input id='sel2' onclick='changeCountry(2)' type='radio' name='optradio'> Country 2</label></div>"
		div.id = "countrySelect";
		this.div = div;
		this.resize();
		return div;
	},
	resize: function() {
		c = $("#map");
		d = $(this.vars.div);
		d.height(c.height() * .1);
		d.width(c.width() * .3)
	}
});
$(document).ready(function(){
	if ($(window).height() < 768 && $(window).width() > 768)
	{
		$("#map").css("height", "100%");
		$("#map").css("width", "50%");
		$("#map").css("float", "left");
		$("#info").css("height", "100%");
		$("#info").css("width", "50%");
		$("#info").css("float", "right");
		map.invalidateSize();

	}
	map = L.map("map", {
		zoom: 1,
		center: [40.702222, -73.979378],
		layers: [cartoLight],
		zoomControl: false,
		attributionControl: false
	});
	$.get("countries.geo.json", function(res){
		allGeojson = res;
		L.geoJSON(allGeojson, {
			onEachFeature: function(f, l){
				l.on("click", function(e){
					console.log(e.target.feature.id);
					selectCountry(e.target.feature.id);
				});
			}
		}).addTo(map);
	});
	cSelect = new L.Control.CountrySelector();
	map.addControl(cSelect);
	L.control.zoom({
		position: "bottomright"
	}).addTo(map);
	changeCountry(1);
	resize();
});
function selectCountry(a3Code) {
	var divString = "#country" + ($("#sel1").is(":checked") ? "One" : "Two");
	$(divString).append("<div id='loader' style='display:hidden'><div class='loader'></div></div>");
	var divEm = $(divString).width() / parseFloat($(divString).css("font-size"));
	$(".loader").css("margin", (divEm - 1) * .25 + "em auto");
	$("#loader").fadeIn("slow", function(e){
		$.get("/country/" + a3Code, function(res){
			
			$("#loader").fadeOut("slow", function(e){
				$("#loader").remove();
			});
		});
	});
}
function changeCountry(num) {
	$("#sel" + num).prop("checked", true);	
	$('.nav li').removeClass('active');
	$("#c" + num + "Head").addClass('active');
	if ($(window).width() < 768){
		if($("#sel1").is(":checked")){
			$("#countryTwo").hide();
			$("#countryOne").show();
		} else {
			$("#countryTwo").show();
			$("#countryOne").hide();
		}
	}
}
function resize(){
	if ($(window).height() < 768 && $(window).width() > 480) {
		$("#map").css("height", "100%");
		$("#map").css("width", "50%");
		$("#map").css("float", "left");
		$("#info").css("height", "100%");
		$("#info").css("width", "50%");
		$("#info").css("float", "right");
		map.invalidateSize();

	} else {
		$("#map").css("height", "50%");
		$("#map").css("width", "100%");
		$("#map").css("float", "none");
		$("#info").css("height", "50%");
		$("#info").css("width", "100%");
		$("#info").css("float", "none");
		map.invalidateSize();

	}
	if ($(window).width() < 768){
		if($("#sel1").is(":checked")){
			$("#countryTwo").show();
			$("#countryOne").hide();
		} else {
			$("#countryTwo").hide();
			$("#countryOne").show();
		}
	} else {
		$("#countryOne, #countryTwo").show();
	}
	$("#countryOne, #countryTwo").css("margin-top", $("#divHead").outerHeight());
	cSelect.resize();
}
$(window).on("resize", resize);
