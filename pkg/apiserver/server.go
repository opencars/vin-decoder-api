package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/opencars/vin-decoder-api/pkg/apiserver/handler"
	"github.com/opencars/vin-decoder-api/pkg/govin"
	"github.com/opencars/vin-decoder-api/pkg/store"
)

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store) *server {
	srv := server{
		router: mux.NewRouter(),
		store:  store,
	}

	srv.configureRouter()

	return &srv
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"X-Api-Key", "Api-Key"})

	cors := handlers.CORS(origins, methods, headers)(s.router)
	cors.ServeHTTP(w, r)
}

func (s *server) decodeVIN() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		lexeme := mux.Vars(r)["vin"]

		if !govin.Valid(lexeme) {
			return handler.ErrInvalidVIN
		}

		vin, err := govin.Parse(lexeme)
		if err != nil {
			return err
		}

		result := NewResult(s.store, vin)
		if err := json.NewEncoder(w).Encode(result); err != nil {
			return err
		}

		return nil
	}
}
