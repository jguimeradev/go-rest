package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jguimeradev/go-rest/internals/store"
)

type Handler struct {
	store *store.Store
}

func New(s *store.Store) *Handler {
	return &Handler{store: s}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.homePage)
	mux.HandleFunc("/health", h.health)
	mux.HandleFunc("GET /users", h.getUsers)
	mux.HandleFunc("POST /users", h.createUser)
	mux.HandleFunc("GET /users/{id}", h.getUser)
	mux.HandleFunc("PUT /users/{id}", h.updateUser)
	mux.HandleFunc("DELETE /users/{id}", h.deleteUser)
}

func (h *Handler) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to SIMPLE REST API")
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	u := h.store.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	u, err := h.store.GetByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var u store.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	created := h.store.Create(u)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(created)
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	var u store.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := r.PathValue("id")
	updated, err := h.store.Update(id, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.store.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
