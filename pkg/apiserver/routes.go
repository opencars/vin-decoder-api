package apiserver

func (s *server) configureRouter() {
	s.router.Handle("/api/v1/vin-decoder/{vin}", s.decodeVIN())
}
