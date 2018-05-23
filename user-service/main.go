package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID           string        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Gender       string        `json:"gender,omitempty"`
	Age          int16         `json:"age,omitempty"`
	UserLocation *UserLocation `json:"userlocation,omitempty"`
}
type UserLocation struct {
	Lat  float32 `json:"Lat,omitempty"`
	Long float32 `json:"Long,omitempty"`
}

var users []User

func main() {
	router := mux.NewRouter().StrictSlash(true)
	users = append(users, User{ID: "1", Name: "Nic", Gender: "Male", Age: 23, UserLocation: &UserLocation{Lat: 54.234, Long: 55.234}})
	router.HandleFunc("/all", GetAllUserEndpoint).Methods("GET")
	router.HandleFunc("/user/{id}", GetUserEndpoint).Methods("GET")
	router.HandleFunc("/create", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/delete/{id}", DeleteUserEndpoint).Methods("DELETE")
	http.ListenAndServe(":30094", router)
}

func GetAllUserEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}

func GetUserEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})

}

func CreateUserEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(users)
}

func DeleteUserEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}
