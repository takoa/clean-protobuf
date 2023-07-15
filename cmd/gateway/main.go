package main

import (
	"log"

	"github.com/takoa/clean-protobuf/internal/app/gateway"
)

func main() {
	if err := gateway.Serve(); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
}
