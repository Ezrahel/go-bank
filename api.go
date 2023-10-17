package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandlerFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(s.handleGetAccount))
	log.Println("JSON API server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAccount(w, r)
	case "POST":
		return s.handleCreateAccount(w, r)
	case "DELETE":
		return s.handleDeleteAccount(w, r)
	default:
		fmt.Errorf("witch cases not aallowed %s", r.Method)
	}
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	WriteJSON(w, http.StatusOK, &Account{})

	account := NewAccount("Israel", "Benkong")
	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type APIServer struct {
	listenAddr string
}

type ApiError struct {
	Error string
}
type apiFunc func(http.ResponseWriter, *http.Request) error

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
