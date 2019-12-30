package apiserver

import (
	"log"
	"net/http"
)

func Start(addr string) error {
	srv := newServer()

	log.Printf("Server is listening on %s...\n", addr)
	return http.ListenAndServe(addr, srv)
}
