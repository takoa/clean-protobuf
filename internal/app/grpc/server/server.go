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

	"github.com/takoa/clean-protobuf/internal/config"
	"github.com/takoa/clean-protobuf/internal/entity/repository"
	routeguidegrpc "github.com/takoa/clean-protobuf/internal/infrastructure/controller/grpc"
	repositoryimpl "github.com/takoa/clean-protobuf/internal/infrastructure/repository"
	routeguidev1 "github.com/takoa/clean-protobuf/internal/pkg/protobuf/routeguide/v1"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
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

	routeGuideServer, err := routeguidegrpc.NewServer(repositories)
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
	routeguidev1.RegisterRouteGuideServiceServer(grpcServer, routeGuideServer)
	reflection.Register(grpcServer)

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
