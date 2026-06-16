//http and store

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handlers struct {
	store *Store
}

func (h *Handlers) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to SIMPLE REST API")
}

func (h *Handlers) health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *Handlers) getUsers(w http.ResponseWriter, r *http.Request) {
	u := h.store.GetAll()
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(u)

}

func (h *Handlers) getUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	u, err := h.store.GetByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func (h *Handlers) createUser(w http.ResponseWriter, r *http.Request) {

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := h.store.Create(newUser)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handlers) updateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser User
	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := r.PathValue("id")
	u, err := h.store.Update(id, updateUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(u)
}
func (h *Handlers) deleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	err := h.store.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

}
