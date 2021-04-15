package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {

}

func RouteInit() *mux.Router{
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,"pong")
	})

	return r
}
