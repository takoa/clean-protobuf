package gateway

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kelseyhightower/envconfig"
	routeguidev1 "github.com/takoa/clean-protobuf/internal/pkg/protobuf/routeguide/v1"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type config struct {
	GRPCServerHost string `default:"localhost" split_words:"true"`
	GRPCServerPort int    `default:"50051" split_words:"true"`
}

var conf config

func init() {
	if err := envconfig.Process("", &conf); err != nil {
		log.Fatalf("Failed to load config: %+v", err)
	}
}

func Serve() error {
	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := routeguidev1.RegisterRouteGuideServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%d", conf.GRPCServerHost, conf.GRPCServerPort), opts)
	if err != nil {
		return xerrors.Errorf("failed to register handlers: %w", err)
	}

	return http.ListenAndServe(":8081", mux)
}
