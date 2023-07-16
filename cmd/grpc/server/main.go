package main

import (
	"log"

	"github.com/takoa/clean-protobuf/internal/app/grpc/server"
)

func main() {
	if err := server.Serve(); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
}
