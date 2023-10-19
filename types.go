package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json"firstname"`
	LastName  string `json"lastname"`
}
type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdat"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(100000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(100000)),
		CreatedAt: time.Now().UTC().AddDate(1, 1, 1),
	}
}
