package grpc

import (
	"context"
	"net"

	"github.com/opencars/grpc/pkg/vin_decoding"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/opencars/vin-decoder-api/pkg/domain"
)

// API represents gRPC API server.
type API struct {
	addr string
	s    *grpc.Server
	svc  domain.InternalService
}

func New(addr string, svc domain.InternalService) *API {
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			RequestLoggingInterceptor,
		),
	}

	return &API{
		addr: addr,
		svc:  svc,
		s:    grpc.NewServer(opts...),
	}
}

// Server returns the underlying gRPC server instance.
func (a *API) Server() *grpc.Server {
	return a.s
}

func (a *API) Run(ctx context.Context) error {
	listener, err := net.Listen("tcp", a.addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	vin_decoding.RegisterServiceServer(a.s, &vinDecodingHandler{api: a})

	// Register reflection service
	reflection.Register(a.s)

	errors := make(chan error)
	go func() {
		errors <- a.s.Serve(listener)
	}()

	select {
	case <-ctx.Done():
		a.s.GracefulStop()
		return <-errors
	case err := <-errors:
		return err
	}
}
