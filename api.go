package main

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func init() {
	users = []User{
		{ID: "1", Username: "User1", Email: "user1@example.com"},
		{ID: "2", Username: "User2", Email: "user2@example.com"},
		{ID: "3", Username: "User3", Email: "user3@example.com"},
		{ID: "4", Username: "User4", Email: "user4@example.com"},
	}
}
