package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Firstname string    `json:"firstName"`
	Lastname  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	LastLogin time.Time `json:"lastLogin"`
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
