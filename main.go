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

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
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
			fmt.Fprintf(w, "User with ID %s deleted", id)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homePage)
	mux.HandleFunc("GET /health", health)
	mux.HandleFunc("GET /users", getUsers)
	mux.HandleFunc("GET /user", getUser)
	mux.HandleFunc("POST /user/create", createUser)
	mux.HandleFunc("PUT /user/update", updateUser)
	mux.HandleFunc("DELETE /user/delete", deleteUser)

	log.Fatal(http.ListenAndServe(":6969", mux))

}
