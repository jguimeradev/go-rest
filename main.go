package main

import (
	"log"
	"net/http"
)

func main() {

	store := NewStore()
	h := &Handlers{store: store}

	mux := http.NewServeMux()
	mux.HandleFunc("/", h.homePage)
	mux.HandleFunc("/health", h.health)
	mux.HandleFunc("GET /users", h.getUsers)
	mux.HandleFunc("POST /users", h.createUser)
	mux.HandleFunc("GET /users/{id}", h.getUser)
	mux.HandleFunc("PUT /users/{id}", h.updateUser)
	mux.HandleFunc("DELETE /users/{id}", h.deleteUser)
	//mux.HandleFunc("POST /user/create", createUser)

	log.Fatal(http.ListenAndServe(":6969", mux))

}
