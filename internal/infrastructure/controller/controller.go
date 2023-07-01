package controller

import (
	"github.com/takoa/clean-protobuf/api"
	"github.com/takoa/clean-protobuf/internal/entity/repository"
	"github.com/takoa/clean-protobuf/internal/usecase/feature"
	"github.com/takoa/clean-protobuf/internal/usecase/route"
	"golang.org/x/xerrors"
)

type RouteGuideServer struct {
	api.UnimplementedRouteGuideServer

	featureFinder      feature.Finder
	routeMessagePoster route.MessagePoster
	routeRecorder      route.InformationGetter
}

func NewServer(repositories *repository.Repositories) (*RouteGuideServer, error) {
	if repositories == nil {
		return nil, xerrors.New("nil repositories")
	}
	return &RouteGuideServer{
		featureFinder:      feature.NewFeatureFinder(repositories.Features),
		routeMessagePoster: route.NewRouteMessagePoster(repositories.Messages),
		routeRecorder:      route.NewRouteInformationGetter(repositories.Features),
	}, nil
}
