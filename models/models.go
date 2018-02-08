package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Country struct {
	ID	bson.ObjectId	`bson:"_id" json:"id"`
	Name	string		`bson:"name" json:"name"`
	LatLng	[] struct {
			Lat	float32
			Lng	float32
		}		`bson:"cover_image" json:"cover_image"`
	Flag	string		`bson:"description" json:"description"`
}
