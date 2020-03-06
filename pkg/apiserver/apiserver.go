package apiserver

import (
	"log"
	"net/http"

	"github.com/opencars/vin-decoder-api/pkg/config"
	"github.com/opencars/vin-decoder-api/pkg/store/sqlstore"
)

// Start starts http server on the specified addr.
func Start(conf *config.Settings, addr string) error {
	store, err := sqlstore.New(&conf.DB)
	if err != nil {
		return err
	}

	srv := newServer(store)

	log.Printf("Server is listening on %s...\n", addr)
	return http.ListenAndServe(addr, srv)
}
