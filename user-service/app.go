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
	params:=mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})

}

func GetAllUserEndpoint(w http.ResponseWriter, req *http.Request){
	json.NewDecoder(w).Encode(users)

}

func CreateUserEndpoint(w http.ResponseWriter, req *http.Request){
	params:=mux.Vars(req)
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)
	users = append(users,user)
	json.NewEncoder(w).Encode(users)
}

func DeleteUserEndpoint(w http.ResponseWriter, req *http.Request){
	params:=mux.Vars(req)
	for index,item  := range people {
		if item.ID == params["id"]{
			users = append(users[:index],users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}
func main(){
	 router:=mux.NewRouter()
	 router.HandleFunc("/allUser",GetUserEndpoint).Methods("GET")
	 router.HandleFunc("/user/{id}",GetAllUserEndpoint).Methods("GET")
	 router.HandleFunc("/createUser",CreateUserEndpoint).Methods("POST")
	 router.HandleFunc("/deleteUser/{id}",DeleteUserEndpoint).Methods("DELETE")
	 log.Fatal(http.ListenAndServe(:8080,router))
}