package apiserver

import (
	"log"
	"net/http"
)

// Start starts http server on the specified addr.
func Start(addr string) error {
	srv := newServer()

	log.Printf("Server is listening on %s...\n", addr)
	return http.ListenAndServe(addr, srv)
}
