package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RouteInit() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	notes := r.PathPrefix("/notes").Subrouter()
	{
		notes.HandleFunc("/", h.GetAllNotes).Methods("GET")
		notes.HandleFunc("/{id}", h.GetByIdNotes).Methods("GET")
		notes.HandleFunc("/", h.StoreNotes).Methods("POST")
		notes.HandleFunc("/{id}", h.StoreNotes).Methods("PUT")
		notes.HandleFunc("/{id}", h.StoreNotes).Methods("DELETE")
	}

	return r
}
