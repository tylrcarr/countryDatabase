package main

import (
	"encoding/json"
	"log"
	"net/http"


	"github.com/gorilla/mux"
	. "github.com/tylrcarr/countryDatabase/config"
	. "github.com/tylrcarr/countryDatabase/dao"
)

var dao = CountriesDb{}
var config = Config{}

// gets all of the countries in one response
func AllCountries(w http.ResponseWriter, r *http.Request) {
	countries, err := dao.FindAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(countries)
}

func FindCountry(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	country, err := dao.FindByName(params["id"])
	if err != nil {
		http.Error(w, "Invalid Country", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(country)
}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/country", AllCountries).Methods("GET")
	r.HandleFunc("/country/{id}", FindCountry).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("site")))
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
