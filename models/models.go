package models

import (
//	"encoding/json"
//	"gopkg.in/mgo.v2/bson"
)
// for getting a specified country
type Country struct {
	Name		string		`bson:"name" json:"name"`
	LatLng		[]float64	`bson:"latlng" json:"latlng"`
	Flag		string		`bson:"flag" json:"flag"`
	Alpha3		string		`bson:"alpha3Code" json:"alpha3"`
	Population	int64		`bson:"population" json:"population"`
	Area		int64		`bson:"area" json:"area"`
}
//for autocomplete of country
type Names struct {
	Names		[]string	`bson:"names" json:"names"`

}
