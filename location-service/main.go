package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	geo "github.com/martinlindhe/google-geolocate"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/location", LocationService).Methods("GET")
	http.ListenAndServe(":2222", router)
}

func LocationService(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := geo.NewGoogleGeo("AIzaSyADmhWTORfjBc-EVYP5RyP3XPC-gzEFZGs")
	res, _ := client.Geolocate()
	fmt.Println(res)
	json.NewEncoder(w).Encode(res)
}
