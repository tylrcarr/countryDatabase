function getCountry()
{
	$.get("/country/" + document.getElementById("country").value, function(res) {
		console.log(res.flag);
		$("#countryGroup").append("<img src='" + res.flag + "' />");
	});
}
