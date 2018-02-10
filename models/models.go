package models

import (
//	"encoding/json"
//	"gopkg.in/mgo.v2/bson"
)
type Properties struct {
	Name	string		`bson:"name" json:"name"`
}
type Polygon struct {
	Lines [][][]float64	`bson:"" json:""`
}
type Geometry struct {
	Coordinates	[][][]float64	`bson:"coordinates" json:"coordinates"`
	Type		string		`bson:"type" json:"type"`
}
type GeoJson struct {
	Geometry	Geometry	`bson:"geometry" json:"geometry"`
	Properties	Properties	`bson:"properties" json:"properties"`
	Type		string		`bson:"type" json:"type"`
}
type Country struct {
//	ID	string		`bson:"name" json:"name"`
	Name	string		`bson:"name" json:"name"`
	LatLng	[]float32	`bson:"latlng" json:"latlng"`
	Flag	string		`bson:"flag" json:"flag"`
	Alpha3	string		`bson:"alpha3Code" json:"alpha3"`
	GeoJson	GeoJson		`bson:"geojson" json:"geojson"`

}
