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
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/takoa/clean-protobuf/api"
	"github.com/takoa/clean-protobuf/internal/config"
	"github.com/takoa/clean-protobuf/internal/entity/model"
	"github.com/takoa/clean-protobuf/internal/entity/repository"
	"github.com/takoa/clean-protobuf/internal/infrastructure/controller"

	repositoryimpl "github.com/takoa/clean-protobuf/internal/infrastructure/repository"
)

func Serve() error {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Server.Port))
	if err != nil {
		return fmt.Errorf("failed to listen %w", err)
	}

	server, err := controller.NewServer(newRepositories())
	if err != nil {
		return fmt.Errorf("failed to initialize the server: %w", err)
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
	api.RegisterRouteGuideServer(grpcServer, server)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

// newRepositories initialize repositories with JSON data.
func newRepositories() *repository.Repositories {
	data, err := os.ReadFile(config.Server.DataFilePath)
	if err != nil {
		log.Fatalf("Failed to load the provided feature data file: %+v", err)
	}

	var features []*model.Feature
	if err := json.Unmarshal(data, &features); err != nil {
		log.Fatalf("Failed to load default features: %+v", err)
	}

	return &repository.Repositories{
		Features: repositoryimpl.NewFeatures(features),
		Messages: repositoryimpl.NewMessages(),
	}
}
