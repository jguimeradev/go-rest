// act as database

package main

import (
	"fmt"
	"strconv"
	"sync"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Store struct {
	mu     sync.RWMutex
	users  map[string]User
	nextID int
}

func NewStore() *Store {
	s := &Store{users: make(map[string]User)}

	// seed data — IDs are clean strings with no trailing spaces (improvement #10)
	seed := []User{
		{ID: "1", Username: "User1", Email: "user1@example.com"},
		{ID: "2", Username: "User2", Email: "user2@example.com"},
		{ID: "3", Username: "User3", Email: "user3@example.com"},
		{ID: "4", Username: "User4", Email: "user4@example.com"},
	}
	for _, u := range seed {
		s.users[u.ID] = u
	}

	// set nextID to the highest existing numeric ID so new IDs never collide (improvement #5)
	for id := range s.users {
		if n, err := strconv.Atoi(id); err == nil && n > s.nextID {
			s.nextID = n
		}
	}

	return s
}

func (s *Store) GetAll() []User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]User, 0, len(s.users))
	for _, u := range s.users {
		result = append(result, u)
	}
	return result
}

func (s *Store) GetByID(id string) (User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	u, ok := s.users[id]
	if !ok {
		return User{}, fmt.Errorf("user %s not found", id)
	}
	return u, nil
}

func (s *Store) Create(u User) User {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nextID++
	u.ID = strconv.Itoa(s.nextID)
	s.users[u.ID] = u
	return u
}

func (s *Store) Update(id string, u User) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.users[id]; !ok {
		return User{}, fmt.Errorf("user %s not found", id)
	}
	u.ID = id
	s.users[id] = u
	return u, nil
}

func (s *Store) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.users[id]; !ok {
		return fmt.Errorf("user %s not found", id)
	}
	delete(s.users, id)
	return nil
}
