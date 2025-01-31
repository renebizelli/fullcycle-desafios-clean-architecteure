package service

import (
	"context"

	"github.com/devfullcycle/20-CleanArch/internal/dtos"
	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrdersUseCase usecase.ListOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := dtos.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.BlankRequest) (*pb.ListOrdersResponse, error) {

	dtos, err := s.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	orders := []*pb.ListOrdersItem{}

	for _, dto := range dtos {
		orders = append(orders, &pb.ListOrdersItem{
			Id:         dto.ID,
			Price:      float32(dto.Price),
			Tax:        float32(dto.Tax),
			FinalPrice: float32(dto.FinalPrice),
		})
	}

	return &pb.ListOrdersResponse{
		Orders: orders,
	}, nil
}
