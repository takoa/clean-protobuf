package main

import (
	"log"

	"github.com/takoa/clean-protobuf/internal/app/connect"
)

func main() {
	if err := connect.Serve(); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
}
