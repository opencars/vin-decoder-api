package main

import (
	"context"
	"flag"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/opencars/seedwork/logger"

	"github.com/opencars/vin-decoder-api/pkg/api/grpc"
	"github.com/opencars/vin-decoder-api/pkg/config"
	"github.com/opencars/vin-decoder-api/pkg/domain/service"
	"github.com/opencars/vin-decoder-api/pkg/store/sqlstore"
)

func main() {
	cfg := flag.String("config", "config/config.yaml", "Path to the configuration file")
	port := flag.Int("port", 3000, "Port of the server")

	flag.Parse()

	conf, err := config.New(*cfg)
	if err != nil {
		logger.Fatalf("config: %v", err)
	}

	logger.NewLogger(logger.LogLevel(conf.Log.Level), conf.Log.Mode == "dev")

	s, err := sqlstore.New(&conf.DB)
	if err != nil {
		logger.Fatalf("store: %v", err)
	}

	r := service.NewInternalService(s.Manufacturer())

	addr := ":" + strconv.Itoa(*port)
	api := grpc.New(addr, r)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger.Infof("Listening on %s...", addr)
	if err := api.Run(ctx); err != nil {
		logger.Fatalf("grpc: %v", err)
	}
}
