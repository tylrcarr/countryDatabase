package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
 
	"github.com/gorilla/mux"
	. "github.com/tylrcarr/countryDatabase/dao"
	. "github.com/tylrcarr/countryDatabase/models"
)

var dao = CountriesDb{"127.0.0.1", "27017"}

func AllCountries(w http.ResponseWriter, r *http.Request) {
	countries, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(countries)
}
 
func FindCountry(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}
 
func CreateCountry(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var country Country
	if err := json.NewDecoder(r.Body).Decode(&country); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	country.ID = bson.NewObjectId()
	if err := dao.Insert(country); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(country)
}

func UpdateCountry(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}
 
func DeleteCountry(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}
 
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/country", AllCountries).Methods("GET")
	r.HandleFunc("/country", CreateCountry).Methods("POST")
	r.HandleFunc("/country", UpdateCountry).Methods("PUT")
	r.HandleFunc("/country", DeleteCountry).Methods("DELETE")
	r.HandleFunc("/country/{id}", FindCountry).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
