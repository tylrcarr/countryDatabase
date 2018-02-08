package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Country struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Name   string        `bson:"name" json:"name"`
	LatLng []float32     `bson:"latlng" json:"latlng"`
	Flag   string        `bson:"flag" json:"flag"`
}
