package dao

import (
	"log"

	. "github.com/tylrcarr/countryDatabase/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CountriesDb struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "data"
)

func (m *CountriesDb) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *CountriesDb) FindAll() ([]Country, error) {
	var countries []Country
	err := db.C(COLLECTION).Find(bson.M{}).All(&countries)
	return countries, err
}

func (m *CountriesDb) FindByName(name string) (Country, error) {
	var country Country
	err := db.C(COLLECTION).Find(bson.M{"name": name}).One(&country)
	return country, err
}

func (m *CountriesDb) Insert(country Country) error {
	err := db.C(COLLECTION).Insert(&country)
	return err
}

func (m *CountriesDb) Delete(country Country) error {
	err := db.C(COLLECTION).Remove(&country)
	return err
}

func (m *CountriesDb) Update(country Country) error {
	err := db.C(COLLECTION).UpdateId(country.ID, &country)
	return err
}
