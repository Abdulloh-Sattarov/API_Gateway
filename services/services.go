package services

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/abdullohsattorov/API_Gateway/config"
	pb "github.com/abdullohsattorov/API_Gateway/genproto/catalog_service"
)

type IServiceManager interface {
	CatalogService() pb.CatalogServiceClient
}

type serviceManager struct {
	catalogService pb.CatalogServiceClient
}

func (s *serviceManager) CatalogService() pb.CatalogServiceClient {
	return s.catalogService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connCatalog, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CatalogServiceHost, conf.CatalogServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		catalogService: pb.NewCatalogServiceClient(connCatalog),
	}

	return serviceManager, nil
}
