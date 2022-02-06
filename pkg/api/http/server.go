package http

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/opencars/seedwork/httputil"
	"github.com/opencars/vin-decoder-api/pkg/domain/command"

	"github.com/opencars/vin-decoder-api/pkg/domain"
	"github.com/opencars/vin-decoder-api/pkg/version"
)

type server struct {
	router *mux.Router

	svc domain.CustomerService
}

func newServer(svc domain.CustomerService) *server {
	srv := server{
		router: mux.NewRouter(),
		svc:    svc,
	}

	srv.configureRouter()

	return &srv
}

func (*server) Version() httputil.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		v := struct {
			Version string `json:"version"`
			Go      string `json:"go"`
		}{
			Version: version.Version,
			Go:      runtime.Version(),
		}

		return json.NewEncoder(w).Encode(v)
	}
}

func (s *server) DecodeVIN() httputil.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		c := command.DecodeVIN{
			UserID:  httputil.UserIDFromContext(r.Context()),
			TokenID: httputil.TokenIDromContext(r.Context()),
			VIN:     mux.Vars(r)["vin"],
		}

		result, err := s.svc.DecodeVIN(r.Context(), &c)
		if err != nil {
			return handleErr(err)
		}

		return json.NewEncoder(w).Encode(result)
	}
}
