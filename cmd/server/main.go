package main

import (
	"log"

	"github.com/opencars/vin-decoder-api/pkg/apiserver"
)

func main() {
	if err := apiserver.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
