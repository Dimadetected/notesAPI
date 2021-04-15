package server

import (
	"net/http"
	"time"
)

type Server struct {
	http *http.Server
}

func Start(port string, handler http.Handler) error {
	server := Server{http: &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
	}}

	err := server.http.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
