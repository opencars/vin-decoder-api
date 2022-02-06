package http

import "github.com/opencars/seedwork/httputil"

func (s *server) configureRouter() {
	v1 := s.router.PathPrefix("/api/v1/").Subrouter()
	v1.Use(
		httputil.CustomerTokenMiddleware(),
	)

	v1.Handle("/vin-decoder/{vin}", s.DecodeVIN())
}
