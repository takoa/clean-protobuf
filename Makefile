.ONESHELL:
SHELL = /bin/bash

NAME = migration

POSTGRES_HOST = localhost
POSTGRES_PORT = 5432
POSTGRES_USER = postgres
POSTGRES_PASSWORD = password
POSTGRES_DATABASE = clean-protobuf
POSTGRES_URL = postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable

PROTOBUF_DIR = ./api
PROTOBUF_FILE = routeguide.proto

KO_DOCKER_REPO = ko.local
SERVER_PATH = ./cmd/server
SERVER_VERSION = v0.0.1
GATEWAY_PATH = ./cmd/gateway
GATEWAY_VERSION = v0.0.1

# Generation
.PHONY: build-images
build-images:
	export KO_DOCKER_REPO=$(KO_DOCKER_REPO)
	export SERVER_VERSION=$(SERVER_VERSION)
	ko build --base-import-paths $(SERVER_PATH)
	export GATEWAY_VERSION=$(GATEWAY_VERSION)
	ko build --base-import-paths $(GATEWAY_PATH)

.PHONY: create-certs
create-certs:
	cd ./internal/pkg/data/x509
	./create.sh

.PHONY: generate-server
generate-server:
	buf mod update $(PROTOBUF_DIR)
	buf generate $(PROTOBUF_DIR)

.PHONY: generate-db-code
generate-db-code:
	go run ./tools/db-code-generator/main.go gorm \
		--dsn "host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DATABASE) sslmode=disable TimeZone=Etc/UTC" \
        --excluded-table schema_migrations \
		--out ./internal/entity/model/generated
	go run ./tools/db-code-generator/main.go model \
		--dsn "host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DATABASE) sslmode=disable TimeZone=Etc/UTC" \
		--excluded-table schema_migrations \
		--out ./internal/entity/model \
		--gorm-model-pkg github.com/takoa/clean-protobuf/internal/entity/model/generated \
		--template-path ./tools/db-code-generator/templates/model.tmpl
	go run ./tools/db-code-generator/main.go repository \
		--dsn "host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DATABASE) sslmode=disable TimeZone=Etc/UTC" \
		--excluded-table schema_migrations \
		--interface-out ./internal/entity/repository \
		--implementation-out ./internal/infrastructure/repository \
		--model-pkg github.com/takoa/clean-protobuf/internal/entity/model \
		--template-path ./tools/db-code-generator/templates/repository.tmpl

# Server Operations
.PHONY: launch-server
launch-server:
	docker compose up -d
	sleep 3

.PHONY: run-client
run-client:
	go run ./cmd/client/main.go

.PHONY: run-server
run-server: build-images launch-server migrate-up-all load-data-into-database

.PHONY: stop-server
stop-server:
	docker compose down

# Database Operations
.PHONY: connect-to-local-db
connect-to-local-db:
	docker compose exec postgres psql $(POSTGRES_URL)

.PHONY: create-new-postgres-migration
create-new-postgres-migration:
	migrate create -ext sql -dir migrations/postgres $(NAME)

.PHONY: load-data-into-database
load-data-into-database:
	go run ./internal/tools/db-data-loader/main.go \
		--dsn "host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DATABASE) sslmode=disable TimeZone=Etc/UTC" \
		--data-file-path "./assets/data.json"

.PHONY: launch-database
launch-database:
	docker compose up -d postgres
	sleep 3

.PHONY: migrate-down-all
migrate-down-all:
	migrate -database $(POSTGRES_URL) -path migrations/postgres down

.PHONY: migrate-down-one
migrate-donw-one:
	migrate -database $(POSTGRES_URL) -path migrations/postgres down 1

.PHONY: migrate-up-all
migrate-up-all:
	migrate -database $(POSTGRES_URL) -path migrations/postgres up

.PHONY: migrate-up-one
migrate-up-one:
	migrate -database $(POSTGRES_URL) -path migrations/postgres up 1

.PHONY: run-database
run-database: launch-database migrate-up-all load-data-into-database

.PHONY: stop-database
stop-database:
	docker compose down

# Tools
# Installing Buf with `go install` isn't a recommended way, but it's easy and universal.
.PHONY: install-tools
install-tools:
	go install github.com/google/ko@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
