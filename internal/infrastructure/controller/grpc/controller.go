package controller

import (
	"github.com/takoa/clean-protobuf/api"
	"github.com/takoa/clean-protobuf/internal/entity/repository"
	"github.com/takoa/clean-protobuf/internal/infrastructure/controller/protobuf"
	"github.com/takoa/clean-protobuf/internal/usecase/feature"
	"github.com/takoa/clean-protobuf/internal/usecase/route"
	"golang.org/x/xerrors"
)

type RouteGuideServer struct {
	api.UnimplementedRouteGuideServer

	getFeatureHandler   *protobuf.GetFeatureHandler
	listFeaturesHandler *protobuf.ListFeaturesHandler
	recordRouteHandler  *protobuf.RecordRouteHandler
	routeChatHandler    *protobuf.RouteChatHandler
}

func NewServer(repositories *repository.Repositories) (*RouteGuideServer, error) {
	if repositories == nil {
		return nil, xerrors.New("nil repositories")
	}
	return &RouteGuideServer{
		getFeatureHandler:   protobuf.NewGetFeatureHandler(feature.NewFeatureFinder(repositories.Features)),
		listFeaturesHandler: protobuf.NewListFeaturesHandler(*feature.NewFeatureFinder(repositories.Features)),
		recordRouteHandler:  protobuf.NewRecordRouteHandler(route.NewRouteInformationGetter(repositories.Features)),
		routeChatHandler:    protobuf.NewRouteChatHandler(route.NewRouteMessagePoster(repositories.Features, repositories.Messages)),
	}, nil
}
