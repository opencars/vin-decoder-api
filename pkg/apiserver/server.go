package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/opencars/vin-decoder-api/pkg/govin"
)

type server struct {
	router *mux.Router
}

func newServer() *server {
	srv := server{
		router: mux.NewRouter(),
	}

	srv.configureRouter()

	return &srv
}

func (s *server) configureRouter() {
	s.router.Handle("/api/v1/vin-decoder/{vin}", s.decodeVIN())
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) decodeVIN() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		lexeme := mux.Vars(r)["vin"]
		if !govin.Valid(lexeme) {
			return ErrInvalidVIN
		}

		vin, err := govin.Parse(mux.Vars(r)["vin"])
		if err != nil {
			return err
		}

		result := NewResult(vin)
		if err := json.NewEncoder(w).Encode(result); err != nil {
			return err
		}

		return nil
	}
}
