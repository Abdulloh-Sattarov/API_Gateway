package services

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/abdullohsattorov/API_Gateway/config"
	pb "github.com/abdullohsattorov/API_Gateway/genproto/catalog_service"
	cl "github.com/abdullohsattorov/API_Gateway/genproto/order_service"
)

type IServiceManager interface {
	CatalogService() pb.CatalogServiceClient
	OrderService() cl.OrderServiceClient
}

type serviceManager struct {
	catalogService pb.CatalogServiceClient
	orderService   cl.OrderServiceClient
}

func (s *serviceManager) CatalogService() pb.CatalogServiceClient {
	return s.catalogService
}

func (s *serviceManager) OrderService() cl.OrderServiceClient {
	return s.orderService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connCatalog, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CatalogServiceHost, conf.CatalogServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connOrder, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.OrderServiceHost, conf.OrderServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		catalogService: pb.NewCatalogServiceClient(connCatalog),
		orderService:   cl.NewOrderServiceClient(connOrder),
	}

	return serviceManager, nil
}
