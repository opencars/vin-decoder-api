package main

import (
	"flag"
	"log"

	"github.com/opencars/vin-decoder-api/pkg/apiserver"
	"github.com/opencars/vin-decoder-api/pkg/config"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "./config/config.yaml", "Path to the configuration file")

	flag.Parse()

	// Get configuration.
	conf, err := config.New(configPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(conf, ":8080"); err != nil {
		log.Fatal(err)
	}
}
