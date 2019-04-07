package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

//Profile Struct (Model)
type Profile struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     *Name  `json:"name"`
}

// Name Struct
type Name struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init users var as a slice Profile struct
var users []Profile

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find with id
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user Profile
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var user Profile
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = params["id"] // Mock ID - not safe
			users = append(users, user)
			json.NewEncoder(w).Encode(users)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)

}

func main() {
	// Init Router
	r := mux.NewRouter()

	//Mock Data
	users = append(users, Profile{ID: "1", Username: "joxi", Email: "johndoe@gmail.com",
		Name: &Name{Firstname: "John", Lastname: "Doe"}})
	users = append(users, Profile{ID: "2", Username: "xclusive1", Email: "pascalx@gmail.com",
		Name: &Name{Firstname: "Pascal", Lastname: "Xclusive"}})

	//Route Handlers / Endpoints
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createProfile).Methods("POST")
	r.HandleFunc("/api/users/{id}", updateProfile).Methods("PUT")
	r.HandleFunc("/api/users/{id}", deleteAccount).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8001", r))
}
