package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
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
func (s *PostgressqlStore) CreateAccount(acct *Account) error {
	query := (`INSERT INTO ACCOUNT 
	(first_name, last_name, number,balance, created_at)
	VALUE($1,$2,$3,$4,$5)`)
	resp, err := s.db.Query(
		query,
		acct.FirstName,
		acct.LastName,
		acct.Number,
		acct.Balance,
		acct.CreatedAt)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
	return nil
}
func (s *PostgressqlStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgressqlStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgressqlStore) GetAccountByID(int) (*Account, error) {
	return nil, nil
}

func (s *PostgressqlStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM ACCOUNT")
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt)

		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
