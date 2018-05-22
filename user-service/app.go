package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

type User struct{
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Gender string `json:"gender,omitempty"`
	Age int16 `json:"age,omitempty"`
	UserLocation *UserLocation `json:"userlocation,omitempty"`
}
type UserLocation struct{
	Lat float32 `json:"Lat,omitempty"`
	Long float32 `json:"Long,omitempty"`
}
var users [] User
func GetUserEndpoint(w http.ResponseWriter, req *http.Request){

}

func GetAllUserEndpoint(w http.ResponseWriter, req *http.Request){
	json.NewDecoder(w).Encode(users)

}

func CreateUserEndpoint(w http.ResponseWriter, req *http.Request){

}

func DeleteUserEndpoint(w http.ResponseWriter, req *http.Request){

}
func main(){
	 router:=mux.NewRouter()
	 router.HandleFunc("/allUser",GetUserEndpoint).Methods("GET")
	 router.HandleFunc("/user/{id}",GetAllUserEndpoint).Methods("GET")
	 router.HandleFunc("/user/{id}",CreateUserEndpoint).Methods("POST")
	 router.HandleFunc("/user/{id}",DeleteUserEndpoint).Methods("DELETE")
	 log.Fatal(http.ListenAndServe(:8080,router))
}