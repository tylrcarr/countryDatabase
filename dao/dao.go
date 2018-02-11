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
	NAMES = "names"
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

func (m *CountriesDb) FindByCode(code string) (Country, error) {
	var country Country
	err := db.C(COLLECTION).Find(bson.M{"alpha3Code": code}).One(&country)
	return country, err
}

func (m *CountriesDb) GetNames() (Names, error) {
	var names Names
	err := db.C(NAMES).Find(bson.M{}).One(&names)
	return names, err
}

func (m *CountriesDb) Insert(country Country) error {
	err := db.C(COLLECTION).Insert(&country)
	return err
}

func (m *CountriesDb) Delete(country Country) error {
	err := db.C(COLLECTION).Remove(&country)
	return err
}

