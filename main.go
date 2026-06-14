package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to SIMPLE REST API")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id := r.URL.Query().Get("id")
	for _, u := range users {
		if u.ID == id {
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-type", "application/json")
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users = append(users, newUser)
	json.NewEncoder(w).Encode(users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	id := r.URL.Query().Get("id")

	var updateUser User

	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, u := range users {
		if u.ID == id {
			users[i] = updateUser
			json.NewEncoder(w).Encode(updateUser)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	for i, u := range users {
		if u.ID == id {
			//slices are passed with spread operator
			users = slices.Delete(users, i, i+1)
			fmt.Fprintf(w, "Book with ID %s deleted", id)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/user", getUser)
	http.HandleFunc("/user/create", createUser)
	http.HandleFunc("/user/update", updateUser)
	http.HandleFunc("/user/delete", deleteUser)

	log.Fatal(http.ListenAndServe(":6969", nil))

}
