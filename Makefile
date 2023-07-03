.ONESHELL:
SHELL = /bin/bash

PROTOBUF_DIR = "./api"
PROTOBUF_FILE = "routeguide.proto"

KO_DOCKER_REPO = ko.local
SERVER_PATH = "./cmd/server"
SERVER_VERSION = "v0.0.1"

.PHONY: generate-server
generate-server:
	cd $(PROTOBUF_DIR)
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	$(PROTOBUF_FILE)

.PHONY: create-certs
create-certs:
	cd "./internal/pkg/data/x509"
	./create.sh

.PHONY: build-images
build-images:
	KO_DOCKER_REPO=$(KO_DOCKER_REPO) SERVER_VERSION=$(SERVER_VERSION) ko build --base-import-paths $(SERVER_PATH)
