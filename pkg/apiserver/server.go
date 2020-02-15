package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
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
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"X-Api-Key", "Api-Key"})

	cors := handlers.CORS(origins, methods, headers)(s.router)
	cors.ServeHTTP(w, r)
}

func (s *server) decodeVIN() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		lexeme := mux.Vars(r)["vin"]
		if !govin.Valid(lexeme) {
			return ErrInvalidVIN
		}

		vin, err := govin.Parse(lexeme)
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
