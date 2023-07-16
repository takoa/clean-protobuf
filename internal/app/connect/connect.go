package connect

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/xerrors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	reflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/takoa/clean-protobuf/api/apiconnect"
	"github.com/takoa/clean-protobuf/internal/config"
	"github.com/takoa/clean-protobuf/internal/entity/repository"
	"github.com/takoa/clean-protobuf/internal/infrastructure/controller/connect"
	repositoryimpl "github.com/takoa/clean-protobuf/internal/infrastructure/repository"
)

func Serve() error {
	flag.Parse()

	repositories, err := newRepositories()
	if err != nil {
		return xerrors.Errorf("failed to create repositories: %w", err)
	}

	routeGuideServer, err := connect.NewServer(repositories)
	if err != nil {
		return xerrors.Errorf("failed to initialize the server: %w", err)
	}

	reflector := reflect.NewStaticReflector("routeguide.RouteGuide")
	reflect.NewHandlerV1(reflector)

	mux := http.NewServeMux()
	mux.Handle(apiconnect.NewRouteGuideHandler(routeGuideServer))
	mux.Handle(reflect.NewHandlerV1(reflector))
	mux.Handle(reflect.NewHandlerV1Alpha(reflector))

	address := fmt.Sprintf(":%d", config.Server.Port)
	log.Printf("server listening at %v", address)
	if err := http.ListenAndServe(
		address,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
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
