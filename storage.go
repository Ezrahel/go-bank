package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) error
}

type PostgressqlStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgressqlStore, error) {
	conStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgressqlStore{
		db: db,
	}, nil
}

func (s *PostgressqlStore) init() error {
	return s.createAccountTable()
}

func (s *PostgressqlStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS ACCOUNT(
		id	serial	primarykey,
		first_name varchar(50),
		last_name varchar(50),
		number	serial,
		balance	serial,
		created_at	timestamp
	)`
	_, err := s.db.Exec(query)
	return err
}
func (s *PostgressqlStore) CreateAccount(*Account) error {
	return nil
}
func (s *PostgressqlStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgressqlStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgressqlStore) GetAccountByID(id int) error {
	return nil
}
