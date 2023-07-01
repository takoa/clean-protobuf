.ONESHELL:
SHELL = /bin/bash

PROTOBUF_DIR = "./api"

.PHONY: generate-server
generate-server:
	cd $(PROTOBUF_DIR)
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	routeguide.proto

.PHONY: create-certs
create-certs:
	cd "./internal/pkg/data/x509"
	./create.sh
