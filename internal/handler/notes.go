package handler

import (
	"fmt"
	"net/http"
)

func (h Handler) GetAllNotes(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"GetAll")
}
func (h Handler) GetByIdNotes(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"GETBYID")
}
func (h Handler) UpdateNotes(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Updated")
}
func (h Handler) DeleteNotes(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"DELETE")
}

func (h Handler) StoreNotes(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"STORE")
}