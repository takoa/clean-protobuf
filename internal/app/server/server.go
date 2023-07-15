/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a simple gRPC server that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It implements the route guide service whose definition can be found in routeguide/route_guide.proto.
package server

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/takoa/clean-protobuf/api"
	"github.com/takoa/clean-protobuf/internal/config"
	"github.com/takoa/clean-protobuf/internal/entity/repository"
	"github.com/takoa/clean-protobuf/internal/infrastructure/controller"
	repositoryimpl "github.com/takoa/clean-protobuf/internal/infrastructure/repository"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func Serve() error {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Server.Port))
	if err != nil {
		return xerrors.Errorf("failed to listen: %w", err)
	}

	repositories, err := newRepositories()
	if err != nil {
		return xerrors.Errorf("failed to create repositories: %w", err)
	}

	healthCheckServer := health.NewServer()
	healthCheckServer.SetServingStatus("RouteGuide", grpc_health_v1.HealthCheckResponse_SERVING)

	routeGuideServer, err := controller.NewServer(repositories)
	if err != nil {
		return xerrors.Errorf("failed to initialize the server: %w", err)
	}

	var opts []grpc.ServerOption
	if config.Server.UsesTLS {
		certFile := config.Server.CertFilePath
		if certFile == "" {
			log.Fatalf("Cert file not provided")
		}
		keyFile := config.Server.KeyFilePath
		if keyFile == "" {
			log.Fatalf("Key file not provided")
		}
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	grpc_health_v1.RegisterHealthServer(grpcServer, healthCheckServer)
	api.RegisterRouteGuideServer(grpcServer, routeGuideServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		return xerrors.Errorf("failed to serve: %w", err)
	}

	return nil
}

func newRepositories() (*repository.Repositories, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.Password,
		config.DB.Name,
		config.DB.SSLMode,
		config.DB.TimeZone,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, xerrors.Errorf("failed to connect to the DB: %w", err)
	}

	return &repository.Repositories{
		Features: repositoryimpl.NewFeatures(db),
		Messages: repositoryimpl.NewMessages(db),
	}, nil
}
