package main

import (
	"context"
	"flag"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/opencars/schema/client"
	"github.com/opencars/seedwork/logger"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/reflection"

	"github.com/opencars/vin-decoder-api/pkg/api/grpc"
	"github.com/opencars/vin-decoder-api/pkg/api/http"
	"github.com/opencars/vin-decoder-api/pkg/config"
	"github.com/opencars/vin-decoder-api/pkg/domain/service"
	"github.com/opencars/vin-decoder-api/pkg/store/sqlstore"
)

func main() {
	cfg := flag.String("config", "config/config.yaml", "Path to the configuration file")
	httpPort := flag.Int("http-port", 8080, "Port for HTTP server")
	grpcPort := flag.Int("grpc-port", 3000, "Port for gRPC server")

	flag.Parse()

	conf, err := config.New(*cfg)
	if err != nil {
		logger.Fatalf("config: %v", err)
	}

	logger.NewLogger(logger.LogLevel(conf.Log.Level), conf.Log.Mode == "dev")

	// Initialize store
	store, err := sqlstore.New(&conf.DB)
	if err != nil {
		logger.Fatalf("store: %v", err)
	}

	// Initialize NATS client
	c, err := client.New(conf.NATS.Address())
	if err != nil {
		logger.Fatalf("nats: %v", err)
	}

	producer, err := c.Producer()
	if err != nil {
		logger.Fatalf("producer: %v", err)
	}

	// Create services
	customerSvc := service.NewCustomerService(store.Manufacturer(), producer)
	internalSvc := service.NewInternalService(store.Manufacturer())

	// Create context with cancellation
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Create errgroup with cancellation
	g, ctx := errgroup.WithContext(ctx)

	// Start HTTP server
	g.Go(func() error {
		addr := ":" + strconv.Itoa(*httpPort)
		logger.Infof("Starting HTTP server on %s...", addr)
		return http.Start(ctx, addr, &conf.Server, customerSvc)
	})

	// Start gRPC server
	g.Go(func() error {
		addr := ":" + strconv.Itoa(*grpcPort)
		logger.Infof("Starting gRPC server on %s...", addr)
		api := grpc.New(addr, internalSvc)

		// Add reflection service
		reflection.Register(api.Server())

		return api.Run(ctx)
	})

	// Wait for interrupt signal or error from servers
	logger.Infof("Servers started successfully. Press Ctrl+C to stop...")
	if err := g.Wait(); err != nil {
		logger.Fatalf("Server error: %v", err)
	}
	logger.Infof("Servers stopped gracefully")
}
